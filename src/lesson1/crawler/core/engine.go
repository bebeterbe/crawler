package core

import (
	"strconv"
	"sync/atomic"
	"time"
)

var (
	logger = NewJournal(Debug)
)

type EngineStatus uint8
type LogType string

const (
	EngineCreated EngineStatus = iota
	EngineInitialized
	EngineStarted
)

//引擎接口
type IEngine interface {
	IClose
	Register(module interface{})
	Start()
	Restart()
}

type Engine struct {
	spider               []ISpider
	pipeline             []IPipeline
	downloader           IDownloader
	downloaderMiddleware []IDownloaderMiddleware
	scheduler            IScheduler

	name        string
	concurrency int32        //最大并发数
	counter     int32        //并发数计数器
	status      EngineStatus //状态
	exit        chan struct{}
}

func (engine *Engine) init() {

	if len(engine.spider) == 0 {
		panic("no usable spider")
	}

	if len(engine.pipeline) == 0 {
		panic("no usable pipeline")
	}

	if engine.scheduler == nil {
		engine.scheduler = ScheduleInstance()
	}

	if engine.downloader == nil {
		engine.downloader = NewDownloader("Downloader")
	}

	for index := range engine.spider {
		go func(i int) {
			for {
				request, hasMore := engine.spider[i].Start()
				engine.scheduler.Push(request)
				if hasMore == false {
					break
				}
			}
		}(index)
	}

	engine.status = EngineInitialized
}

func (engine *Engine) handle(request *Request) {

	atomic.AddInt32(&(engine.counter), 1)

	defer func() {
		atomic.AddInt32(&(engine.counter), -1)
	}()

	downloader := engine.downloader

	response, err := downloader.Start(request, engine.downloaderMiddleware)

	if err != nil {
		logger.Error(err)
		return
	}

	if response == nil {
		return
	}

	parser := response.Request.Parser

	if parser == nil {
		return
	}

	spider := response.Request.Spider

	//将返回转交给解析器进行解析
	requests, items, err := parser(spider, response)

	if err != nil {
		logger.Error(err)
		return
	}

	//将解析后的结果交给调度器
	if len(requests) > 0 {
		engine.scheduler.BatchPush(requests)
	}

	//将解析后的结果交给pipeline处理
	if len(items) > 0 {
		current := items
		for index := range engine.pipeline {
			current, err = engine.pipeline[index].Run(current)
			if err != nil {
				logger.Error(err)
				break
			}
		}
	}
}

func (engine *Engine) dispatch() {
	for {
		if engine.concurrency < engine.counter {
			time.Sleep(time.Second)
			continue
		}

		request := engine.scheduler.Pop()

		if request == nil {
			time.Sleep(time.Second)
			continue
		}

		go engine.handle(request)
	}
}

//注册组价
func (engine *Engine) Register(module interface{}) {
	if t, ok := module.(ISpider); ok {
		engine.spider = append(engine.spider, t)
		return
	}

	if t, ok := module.(IDownloader); ok {
		engine.downloader = t
		return
	}

	if t, ok := module.(IPipeline); ok {
		engine.pipeline = append(engine.pipeline, t)
		return
	}

	if t, ok := module.(IScheduler); ok {
		engine.scheduler = t
		return
	}

	if t, ok := module.(IDownloaderMiddleware); ok {
		engine.downloaderMiddleware = append(engine.downloaderMiddleware, t)
		return
	}

	panic("not support this kind of module")
}

func (engine *Engine) Name() string {
	return engine.name
}

func (engine *Engine) String() string {
	return "status:" + string(engine.status) +
		", max concurrency:" + strconv.Itoa(int(engine.concurrency)) +
		", current concurrency:" + strconv.Itoa(int(engine.counter))
}

func (engine *Engine) Start() {
	engine.init()

	go engine.dispatch()

	engine.status = EngineStarted

	<-engine.exit
}

func (engine *Engine) Restart() {

}

func (engine *Engine) Close() error {
	return nil
}

func NewEngine(concurrency int32) *Engine {

	if concurrency < 1 {
		concurrency = 1
	}

	return &Engine{
		name:        "Engine",
		status:      EngineCreated,
		concurrency: concurrency,
		exit:        make(chan struct{}, 1),
	}
}
