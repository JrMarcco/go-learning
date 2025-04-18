package main

import (
	"github.com/JrMarcco/go-learning/fx/ioc"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		ioc.AppFxOpt,
		ioc.RouteFxOpt,
		ioc.ZapFxOpt,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()

}
