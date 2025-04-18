package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServerMux,
				//fx.ParamTags(`name:"echoHandler"`, `name:"helloHandler"`),
				fx.ParamTags(`group:"routes"`),
			),
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
				//fx.ResultTags(`name:"echoHandler"`),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				NewHelloHandler,
				fx.As(new(Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				NewEchoService,
				fx.As(new(EchoService)),
			),
			zap.NewExample,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(func(*http.Server) {}),
	).Run()

}

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

type Route interface {
	http.Handler
	Pattern() string
}

type EchoHandler struct {
	logger *zap.Logger
	echo   EchoService
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.logger.Warn("Failed to copy request body", zap.Error(err))
	}
	msg, _ := h.echo.Echo(r.Context())
	h.logger.Info(msg)
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}

//func NewServerMux(echo, hello Route) *http.ServeMux {
//	mux := http.NewServeMux()
//	mux.Handle(echo.Pattern(), echo)
//	mux.Handle(hello.Pattern(), hello)
//	return mux
//}

func NewServerMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

type HelloHandler struct {
	logger *zap.Logger
}

func (h *HelloHandler) ServeHTTP(w_ http.ResponseWriter, _ *http.Request) {
	h.logger.Info("hello")
}

func (h *HelloHandler) Pattern() string {
	return "/hello"
}

func NewHelloHandler(logger *zap.Logger) *HelloHandler {
	return &HelloHandler{
		logger: logger,
	}
}

type EchoService interface {
	Echo(ctx context.Context) (string, error)
}

type EchoServiceImpl struct{}

func (e *EchoServiceImpl) Echo(_ context.Context) (string, error) {
	return "hello world", nil
}

func NewEchoService() *EchoServiceImpl {
	return &EchoServiceImpl{}
}

func NewEchoHandler(logger *zap.Logger, echoService EchoService) *EchoHandler {
	return &EchoHandler{
		logger: logger,
		echo:   echoService,
	}
}
