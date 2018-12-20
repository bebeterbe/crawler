package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Method string

const (
	GET  Method = "GET"
	PUT  Method = "PUT"
	DEL  Method = "DELETE"
	POST Method = "POST"
	HEAD Method = "HEAD"
)

//请求参数
type Query map[string][]string
type Header map[string][]string
type Body interface{}

//不同类型的body
type JsonBody []byte
type XmlBody []byte
type FormBody url.Values

type Request struct {
	*http.Request

	proxy     *url.URL
	timeout   time.Duration
	cookieJar http.CookieJar

	Parser     Parser      //解析函数
	Spider     ISpider     //爬虫定义
	Downloader IDownloader //下载器
}

func (request *Request) String() string {
	return string(request.Method) + " " + request.URL.String()
}

func (request *Request) SetBody(body interface{}) (*Request, error) {
	switch t := body.(type) {
	case JsonBody:
		{
			request.Request.Body = ioutil.NopCloser(bytes.NewReader(t))
			request.Header.Add("Content-Type", "application/json; charset=utf-8")
		}
	case XmlBody:
		{
			request.Request.Body = ioutil.NopCloser(bytes.NewReader(t))
			request.Header.Add("Content-Type", "application/xml; charset=UTF-8")
		}
	case FormBody:
		{
			request.Request.Body = ioutil.NopCloser(strings.NewReader(url.Values(t).Encode()))
			request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		}
	case string:
		{
			request.Request.Body = ioutil.NopCloser(strings.NewReader(t))
			request.Header.Add("Content-Type", "text/plain; charset=utf-8")
		}
	case []byte:
		{
			request.Request.Body = ioutil.NopCloser(bytes.NewReader(t))
		}
	case bytes.Buffer:
		{
			request.Request.Body = ioutil.NopCloser(bytes.NewReader(t.Bytes()))
		}
	case io.ReadCloser:
		{
			request.Request.Body = t
		}
	default:
		return request, errors.New(fmt.Sprintf("not support this type '%v' for request body", t))
	}

	return request, nil
}

func (request *Request) SetQuery(query Query) (*Request, error) {
	for key, value := range query {
		request.URL.Query().Add(key, strings.Join(value, ","))
	}
	return request, nil
}

func (request *Request) SetProxy(uri string) (*Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return request, err
	}
	request.proxy = u
	return request, nil
}

func (request *Request) SetTimeout(d time.Duration) (*Request, error) {
	request.timeout = d
	return request, nil
}

func (request *Request) SetCookieJar(jar http.CookieJar) (*Request, error) {
	request.cookieJar = jar
	return request, nil
}

func (request *Request) GetProxy() string {
	if request.proxy == nil {
		return ""
	}

	return request.proxy.String()
}

func (request *Request) GetTimeout() time.Duration {
	return request.timeout
}

func (request *Request) GetCookieJar() http.CookieJar {
	return request.cookieJar
}

func NewRequest(uri string, method Method, parser Parser, spider ISpider) (*Request, error) {

	//注意这里method必须为大写
	request, err := http.NewRequest(string(method), uri, nil)

	if err != nil {
		return nil, err
	}

	return &Request{
		Request: request,
		Parser:  parser, //解析函数
		Spider:  spider,
		timeout: -1,
		proxy:   nil,
	}, nil
}
