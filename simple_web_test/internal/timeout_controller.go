package internal

import (
	"github.com/jrmarcco/go-learning/simple_web/framework"
	"time"
)

func TimeoutController(ctx *framework.Context) {
	time.Sleep(10 * time.Second)
	ctx.SetOkStatus().Json("ok, TimeoutController")
}
