package huxiu

import (
	"lesson2/crawler/core"
	"log"
)

type Pipeline struct {
	core.Pipeline
}

func (pipe *Pipeline) Close() error {
	return nil
}

func (pipe *Pipeline) Run(items []core.Item) ([]core.Item, error) {
	log.Println(items)
	return nil, nil
}
