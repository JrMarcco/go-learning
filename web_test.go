package go_learning

import (
	"github.com/jrmarcco/go-learning/web"
	"github.com/jrmarcco/go-learning/web/middleware/accesslog"
	"log"
	"testing"
)

func TestWeb_HttpServer(t *testing.T) {

	h := web.NewHttpServer(":8080")

	userApi := h.Group("/user")
	userApi.Use(accesslog.NewBuilder().Build())

	{
		userApi.Get("/list", func(ctx *web.Context) {
			if err := ctx.OkJson("hello"); err != nil {
				log.Fatalln(err)
			}
		})
	}

	_ = h.Start()

}
