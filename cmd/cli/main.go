package main

import (
	"os"
	root "scheduler/internal/cli"
)

func main() {

	if err := root.NewRootCommand().Register().Execute(); err != nil {
		os.Exit(0)
	}
}
