package main

import (
	"net/http"

	"github.com/JrMarcco/go-learning/fx/ioc"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		ioc.ZapFxOpt,
		ioc.AppFxOpt,
		ioc.RouteFxOpt,

		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()

}
