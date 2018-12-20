package huxiu

import (
	"crawler/core"
)

type Spider struct {
	core.Spider
}

func (spider *Spider) Close() error {
	return nil
}

func (spider *Spider) Parser(s core.ISpider, response *core.Response) ([]*core.Request, []core.Item, error) {
	var items []core.Item
	items = append(items, 1)
	return nil, items, nil
}

func (spider *Spider) Start() (*core.Request, bool) {
	request, err := core.NewRequest("https://www.huxiu.com/article/182059.html", core.GET, spider.Parser, spider)

	if err != nil {
		return nil, false
	}

	return request, false
}
