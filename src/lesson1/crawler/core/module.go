package core

import (
	"fmt"
)

type IClose interface {
	Close() error //用于清除资源
}

type IModule interface {
	IClose
	fmt.Stringer
	Name() string
}
