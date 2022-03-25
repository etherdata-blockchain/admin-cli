package main

import (
	"fmt"
	"os"

	"cli/internal/engine"
)

func main() {
	//etdClient := clients.ETD{}
	engine := engine.NewFlagEngine()

	if err := engine.Init(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if err := engine.Run(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
