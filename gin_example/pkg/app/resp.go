package app

import (
	"github.com/gin-gonic/gin"
	"go_learning/gin_example/pkg/e"
)

type Gin struct {
	Ctx *gin.Context
}

func (g *Gin) Resp(httpCode, errCode int, data any) {
	g.Ctx.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
