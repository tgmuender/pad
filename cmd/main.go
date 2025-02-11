package main

import (
	"fmt"
	"os"
	"xgmdr.com/pad/cmd/cmd"
)

// main is the entry point of the application
func main() {
	var cmd = cmd.NewRootCmd()
	err := cmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
