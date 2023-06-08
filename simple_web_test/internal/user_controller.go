package internal

import "github.com/jrmarcco/go-learning/simple_web/framework"

func UserController(ctx *framework.Context) {
	ctx.SetOkStatus().Json("ok, UserController")
}
