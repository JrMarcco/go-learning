package ioc

import (
	"github.com/JrMarcco/go-learning/fx/route"
	"go.uber.org/fx"
)

var RouteFxOpt = fx.Provide(
	fx.Annotate(
		route.NewEchoHandler,
		fx.As(new(route.Route)),
		//fx.ResultTags(`name:"echoHandler"`),
		fx.ResultTags(`group:"routes"`),
	),
	fx.Annotate(
		route.NewHelloHandler,
		fx.As(new(route.Route)),
		//fx.ResultTags(`name:"helloHandler"`),
		fx.ResultTags(`group:"routes"`),
	),
)
