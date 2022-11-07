package jwt

import (
	"github.com/gin-gonic/gin"
	"go_learning/gin_example/pkg/e"
	"go_learning/gin_example/pkg/util"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.SUCCESS
		toke := ctx.Query("token")
		if toke == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(toke)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
