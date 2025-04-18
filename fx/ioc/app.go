package ioc

import (
	"context"
	"github.com/JrMarcco/go-learning/fx/route"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
)

var AppFxOpt = fx.Provide(
	NewHTTPServer,
	fx.Annotate(
		NewServerMux,
		//fx.ParamTags(`name:"echoHandler"`, `name:"helloHandler"`),
		fx.ParamTags(`group:"routes"`),
	),
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, logger *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			logger.Info("HTTP server listening", zap.String("addr", srv.Addr))
			go func() {
				_ = srv.Serve(ln)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

//func NewServerMux(echo, hello Route) *http.ServeMux {
//	mux := http.NewServeMux()
//	mux.Handle(echo.Pattern(), echo)
//	mux.Handle(hello.Pattern(), hello)
//	return mux
//}

func NewServerMux(routes []route.Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.Pattern(), r)
	}
	return mux
}
