package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sidechannelinc/enclave-code-challenge/pkg/efw"
)

const defaultErrorMessage string = "expected 'status' or 'sync' subcommands"

func main() {
	e := efw.New(context.Background())

	if len(os.Args) < 2 {
		fmt.Println(defaultErrorMessage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "status":
		e.Status()
	case "sync":
		e.Sync()
	case "help":
		// TODO
		// implement a help command
		fmt.Println("TODO implement a help command")
	default:
		fmt.Println(defaultErrorMessage)
		os.Exit(1)
	}
}
