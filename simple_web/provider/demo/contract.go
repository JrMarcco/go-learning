package demo

const Key = "simple_web:provider:demo"

type FooService interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
