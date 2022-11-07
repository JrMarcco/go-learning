package demo

import (
	"fmt"
	"go_learning/simple_web/framework"
)

type ServiceDemo struct {
	FooService
	container framework.Container
}

func NewServiceDemo(params ...any) (any, error) {
	container := params[0].(framework.Container)

	fmt.Println("new service demo")

	return &ServiceDemo{
		container: container,
	}, nil
}

func (s *ServiceDemo) GetFoo() Foo {
	return Foo{
		Name: "Foo",
	}
}
