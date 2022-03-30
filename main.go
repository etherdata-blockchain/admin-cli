package main

import (
	"fmt"
	"os"

	"cli/internal/engine"
)

func main() {
	//etdClient := clients.ETD{}
	selectedEngine := engine.NewFlagEngine()

	if err := selectedEngine.Init(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if err := selectedEngine.Run(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
