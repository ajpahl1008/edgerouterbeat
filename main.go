package main

import (
	"github.com/ajpahl1008/edgerouterbeat/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
