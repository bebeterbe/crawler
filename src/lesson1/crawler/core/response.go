package core

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Response struct {
	*http.Response
	Request *Request
}

func (res *Response) Bytes() ([]byte, error) {

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (res *Response) String() string {
	bytes, _ := res.Bytes()
	return string(bytes)
}

func (res *Response) Json(value interface{}) error {
	bytes, err := res.Bytes()

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, value)
}

func (res *Response) Xml(value interface{}) error {
	bytes, err := res.Bytes()

	if err != nil {
		return err
	}

	return xml.Unmarshal(bytes, value)
}

func (res *Response) File(p string) error {
	defer res.Body.Close()

	dir := path.Base(p)

	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	file, err := os.Create(p)

	if err != nil {
		return err
	}

	defer file.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(file, res.Body)

	return err
}

func NewResponse(req *Request, res *http.Response) *Response {
	return &Response{
		Request:  req,
		Response: res,
	}
}
