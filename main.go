package main

import (
	"log"

	"go.sbr.pm/lord/cli"

	// Enable builders
	_ "go.sbr.pm/lord/builders/golang"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
