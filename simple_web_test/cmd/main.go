package main

import (
	"context"
	"go_learning/simple_web/framework"
	"go_learning/simple_web/framework/middleware"
	"go_learning/simple_web/provider/demo"
	"go_learning/simple_web_test/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	internal.RegisterRouter(core)

	if err := core.Bind(&demo.ServiceProviderDemo{}); err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server Start: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}
