package internal

import (
	"go_learning/simple_web/framework"
	"time"
)

func TimeoutController(ctx *framework.Context) {
	time.Sleep(10 * time.Second)
	ctx.SetOkStatus().Json("ok, TimeoutController")
}
