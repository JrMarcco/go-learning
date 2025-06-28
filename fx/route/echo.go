package route

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type EchoHandler struct {
	logger *zap.Logger
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.logger.Warn("Failed to copy request body", zap.Error(err))
	}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}

func NewEchoHandler(logger *zap.Logger) *EchoHandler {
	return &EchoHandler{
		logger: logger,
	}
}
