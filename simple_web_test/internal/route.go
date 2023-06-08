package internal

import (
	"github.com/jrmarcco/go-learning/simple_web/framework"
	"github.com/jrmarcco/go-learning/simple_web/framework/middleware"
)

func RegisterRouter(core *framework.Core) {
	core.Get("/timeout", TimeoutController)
	core.Get("/user/login", middleware.Cost(), UserController)

	subApi := core.Group("/sub")
	{
		subApi.Use(middleware.Cost())

		subApi.Get("/:id", SubjectGetController)
		subApi.Post("/:id", SubjectAddController)
		subApi.Put("/:id", SubjectUpdateController)
		subApi.Delete("/:id", SubjectDelController)
		subApi.Get("/list/all", SubjectListController)

		subInnerApi := subApi.Group("/info")
		{
			subInnerApi.Get("/name", SubjectNameController)
		}
	}
}
