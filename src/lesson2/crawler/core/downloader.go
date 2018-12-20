package core

import (
	"errors"
	"net/http"
	"strings"
)

type IDownloaderMiddleware interface {
	ProcessRequest(request *Request) error
	ProcessResponse(response *Response) error
}

type IDownloader interface {
	IModule
	Start(request *Request, middleware []IDownloaderMiddleware) (*Response, error)
}

type Downloader struct {
	stop bool
	name string
}

func (downloader *Downloader) Name() string {
	return downloader.name
}

func (downloader *Downloader) String() string {
	return downloader.name
}

func (downloader *Downloader) Close() error {
	downloader.stop = true
	return nil
}

func (downloader *Downloader) Start(request *Request, middleware []IDownloaderMiddleware) (*Response, error) {
	if downloader.stop == true {
		return nil, errors.New("engine has stop")
	}

	request.Method = strings.ToUpper(request.Method)

	httpClient := &http.Client{}

	if request.cookieJar != nil {
		httpClient.Jar = request.cookieJar
	}

	if request.timeout > 0 {
		httpClient.Timeout = request.timeout
	}

	if request.proxy != nil {
		httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(request.proxy),
		}
	}

	request.Downloader = downloader

	for index := range middleware {
		err := middleware[index].ProcessRequest(request)

		if err != nil {
			return nil, err
		}
	}

	res, err := httpClient.Do(request.Request)

	if err != nil {
		return nil, err
	}

	response := NewResponse(request, res)

	for index := range middleware {
		middleware[index].ProcessResponse(response)

		if err != nil {
			return response, err
		}
	}

	return response, nil

}

func NewDownloader(name string) *Downloader {
	return &Downloader{
		name: name,
		stop: false,
	}
}
