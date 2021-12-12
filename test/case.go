package test

import (
	"context"
)

type Case struct {
	Name     string
	Setup    func() interface{}
	Ctx      context.Context
	Args     interface{}
	Expected interface{}
	IsErr    bool
	Err      interface{}
}
