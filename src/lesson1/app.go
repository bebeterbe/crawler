package lesson1

import (
	"lesson1/crawler/core"
	"lesson1/custom"
)

func main() {
	engine := core.NewEngine(10)

	spider := &huxiu.Spider{}
	pipeline := &huxiu.Pipeline{}
	middleware := &huxiu.HeaderMiddleware{}

	engine.Register(spider)
	engine.Register(pipeline)
	engine.Register(middleware)

	engine.Start()
}
