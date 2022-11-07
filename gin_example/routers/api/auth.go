package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_learning/gin_example/models"
	"go_learning/gin_example/pkg/e"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	valid := validation.Validation{}

	data := make(map[string]any)
	code := e.INVALID_PARAMS
	if ok, _ := valid.Valid(&auth{Username: username, Password: password}); ok {
		if exist := models.CheckAuth(username, password); exist {
			token, err := util.GenerateToken(username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
