package main

import (
	"fmt"
	"os"

	"cli/internal/clients"
	"cli/internal/engine"
)

func main() {
	etdClient := &clients.ETD{}
	zipClient := &clients.Zip{}

	selectedEngine := engine.NewFlagEngine(etdClient, zipClient)

	if err := selectedEngine.Init(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if err := selectedEngine.Run(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
