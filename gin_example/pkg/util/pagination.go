package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_learning/gin_example/pkg/setting"
)

func GetPage(c *gin.Context) int {

	if page, _ := com.StrTo(c.Query("page")).Int(); page > 0 {
		return (page - 1) * setting.AppSetting.PageSize
	}

	return 0
}
