package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/url"
	"strconv"
)

func main() {
	engine := gin.New()

	// counter
	engine.GET("/counter", func(ctx *gin.Context) {
		purl, _ := url.Parse(ctx.Request.RequestURI)
		AccessCounter.With(prometheus.Labels{
			"method": ctx.Request.Method,
			"path":   purl.Path,
		}).Add(1)
	})

	// gauge
	engine.GET("queue", func(ctx *gin.Context) {
		num := ctx.Query("num")
		fnum, _ := strconv.ParseFloat(num, 32)
		QueueGauge.With(prometheus.Labels{
			"name": "queue_demo",
		}).Set(fnum)
	})

	// histogram
	engine.GET("/histogram", func(ctx *gin.Context) {
		purl, _ := url.Parse(ctx.Request.RequestURI)
		HttpDurationHistogram.With(
			prometheus.Labels{
				"path": purl.Path,
			},
		).Observe(float64(rand.Intn(30)))
	})

	// summary
	engine.GET("/summary", func(ctx *gin.Context) {
		purl, _ := url.Parse(ctx.Request.RequestURI)
		HttpDurations.With(prometheus.Labels{
			"path": purl.Path,
		}).Observe(float64(rand.Intn(30)))
	})

	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
