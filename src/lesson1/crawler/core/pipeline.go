package core

type Item interface{}

type IPipeline interface {
	IModule
	Run(item []Item) ([]Item, error)
}

type Pipeline struct {
	name string
}

func (pipe *Pipeline) Name() string {
	return pipe.name
}

func (pipe *Pipeline) String() string {
	return pipe.name
}

func (pipe *Pipeline) Close() error {
	panic("this method has not implement")
}

func (pipe *Pipeline) Run(item []Item) ([]Item, error) {
	panic("this method has not implement")
}
