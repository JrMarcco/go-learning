package web

import (
	"net"
	"net/http"
)

type Server interface {
	http.Handler

	Start() error
}

type HttpServer struct {
	*router

	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr:   addr,
		router: newRouter(),
	}
}

var _ Server = &HttpServer{}

func (h *HttpServer) Start() error {
	l, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}

	return http.Serve(l, h)
}

func (h *HttpServer) Get(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodGet, path, handleFunc)
}

func (h *HttpServer) Post(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodPost, path, handleFunc)
}

func (h *HttpServer) Put(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodPut, path, handleFunc)
}

func (h *HttpServer) Delete(path string, handleFunc HandleFunc) {
	h.addRoute(http.MethodDelete, path, handleFunc)
}

func (h *HttpServer) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {

	ctx := &Context{
		req: req,
		rsp: rsp,
	}

	h.serve(ctx)
}

func (h *HttpServer) serve(ctx *Context) {

	mi := h.findRoute(ctx.req.Method, ctx.req.URL.Path)
	if mi.matched {
		mi.handleFunc(ctx)
		return
	}

	ctx.rsp.WriteHeader(http.StatusNotFound)
	_, _ = ctx.rsp.Write([]byte("Not Found"))
}
