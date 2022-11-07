package main

import (
	"go_learning/word-cli/cmd"
	"log"
)

func main() {
	err := cmd.Exec()
	if err != nil {
		log.Fatalf("Fail to execute command: %v\n", err)
	}
}
