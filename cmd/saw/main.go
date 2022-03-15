package main

import (
	"os"
)

func main() {
	if err := sawCommand.Execute(); err != nil {
		os.Exit(0)
	}
}
