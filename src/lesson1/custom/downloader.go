package huxiu

import (
	"lesson1/crawler/core"
)

type HeaderMiddleware struct {
}

func (middleware *HeaderMiddleware) ProcessRequest(request *core.Request) error {
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("o", "keep-alive")
	request.Header.Add("Pragma", "no-cache")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	return nil
}

func (middleware *HeaderMiddleware) ProcessResponse(response *core.Response) error {
	return nil
}
