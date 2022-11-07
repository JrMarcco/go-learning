package middleware

import (
	"go_learning/simple_web/framework"
	"log"
	"time"
)

func Cost() framework.HandlerFunc {
	return func(ctx *framework.Context) {
		start := time.Now()

		ctx.Next()

		log.Printf("api uri: %v, cost: %v", ctx.Request().RequestURI, time.Now().Sub(start).Seconds())
	}
}
