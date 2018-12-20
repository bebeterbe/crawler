package core

import (
	"log"
	"sync"
)

var (
	once      sync.Once
	scheduler *Scheduler
)

//调度器接口
type IScheduler interface {
	IModule
	Pop() *Request
	Push(request *Request)
	BatchPush(request []*Request)
}

//主调度器
type Scheduler struct {
	request []*Request
	name    string
}

func (scheduler *Scheduler) Name() string {
	return scheduler.name
}

func (scheduler *Scheduler) String() string {
	return scheduler.name
}

func (scheduler *Scheduler) Close() error {
	return nil
}

func (scheduler *Scheduler) Pop() *Request {
	//该方法会被多个协程调用，注意并发安全
	if len(scheduler.request) == 0 {
		return nil
	}

	request := scheduler.request[0]
	scheduler.request = scheduler.request[1:]

	return request
}

func (scheduler *Scheduler) Push(request *Request) {
	scheduler.request = append(scheduler.request, request)
}

func (scheduler *Scheduler) BatchPush(requests []*Request) {
	for _, request := range requests {
		scheduler.request = append(scheduler.request, request)
	}

	log.Println(scheduler.request)
}

func ScheduleInstance() *Scheduler {
	if scheduler != nil {
		return scheduler
	}

	once.Do(func() {
		scheduler = &Scheduler{
			name: "Scheduler",
		}
	})

	return scheduler
}
