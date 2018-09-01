package main

import (
	"os"

	"github.com/ajpahl1008/edgerouterbeat/cmd"

	_ "github.com/ajpahl1008/edgerouterbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
