package core

var (
	spiderLogger = NewJournal(Info)
)

type ISpider interface {
	IModule
	Start() (*Request, bool)
}

type ParserResult struct {
	Requests []*Request
	Items    []Item
}

type Parser func(spider ISpider, response *Response) ([]*Request, []Item, error)

type Spider struct {
	name string
}

func (spider *Spider) Name() string {
	return spider.name
}

func (spider *Spider) String() string {
	return spider.name
}

func (spider *Spider) Log(level Level, message string, args ...interface{}) {
	spiderLogger.Log(level, message, args)
}

func (spider *Spider) Close() error {
	panic("this method has not implement")
}

func (spider *Spider) Start() (request *Request, hasMore bool) {
	panic("this method has not implement")
}

func NewSpider(name string) *Spider {
	return &Spider{
		name: name,
	}
}
