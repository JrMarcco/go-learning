package web

import (
	"log"
	"testing"
)

func TestHttpServer(t *testing.T) {

	h := NewHttpServer(":8080")

	h.Get("/test", func(ctx *Context) {
		log.Println("Hello World")
	})

	_ = h.Start()
}
