package main

import (
	"github.com/robfig/cron"
	"go_learning/gin_example/models"
	"log"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	log.Println("Starting ...")

	c := cron.New()
	err := c.AddFunc("* * * * * ?", func() {
		log.Println("Run model.CleanAllTag ...")
		models.CleanAllTag()
	})
	if err != nil {
		log.Fatalln(err)
	}

	c.Start()

	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			timer.Reset(10 * time.Second)
		}
	}
}
