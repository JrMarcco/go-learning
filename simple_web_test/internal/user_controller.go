package internal

import "go_learning/simple_web/framework"

func UserController(ctx *framework.Context) {
	ctx.SetOkStatus().Json("ok, UserController")
}
