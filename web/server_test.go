package web

import (
	"log"
	"testing"
)

func TestHttpServer(t *testing.T) {

	h := NewHttpServer(":8080")

	h.Get("/user", func(context *Context) {
		log.Println("user")
	})

	userApi := h.Group("/user")
	userApi.Get("/get", func(context *Context) {
		log.Println("user get")
	})
	userApi.Post("/post", func(context *Context) {
		log.Println("user post")

	})

	delUserApi := userApi.Group("/delete")
	delUserApi.Get("/all", func(context *Context) {
		log.Println("delete all")
	})

	_ = h.Start()
}
