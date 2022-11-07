package app

import (
	"github.com/astaxie/beego/validation"
	"go_learning/gin_example/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
