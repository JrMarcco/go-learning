package main

import (
	"log"

	"github.com/JrMarcco/go-learning/word-cli/cmd"
)

func main() {
	err := cmd.Exec()
	if err != nil {
		log.Fatalf("Fail to execute command: %v\n", err)
	}
}
