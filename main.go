package main

import (
	"log"

	"go.sbr.pm/lord/cli"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
