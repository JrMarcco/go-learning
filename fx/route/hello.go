package route

import (
	"go.uber.org/zap"
	"net/http"
)

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
