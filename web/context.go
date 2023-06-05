package web

import "net/http"

type Context struct {
	req *http.Request
	rsp http.ResponseWriter
}

type HandleFunc func(ctx *Context)
