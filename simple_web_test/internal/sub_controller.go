package internal

import (
	"github.com/jrmarcco/go-learning/simple_web/framework"
	"github.com/jrmarcco/go-learning/simple_web/provider/demo"
)

func SubjectAddController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectAddController")
}

func SubjectListController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectListController")
}

func SubjectDelController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectDelController")
}

func SubjectUpdateController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
}

func SubjectGetController(c *framework.Context) {
	c.SetOkStatus().Json("ok, SubjectGetController")
}

func SubjectNameController(c *framework.Context) {

	serviceDemo := c.MustMake(demo.Key).(demo.FooService)
	foo := serviceDemo.GetFoo()

	c.SetOkStatus().Json(foo)
}
