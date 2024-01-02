package main

import (
	"fmt"
	"os"
	"xgmdr.com/pad/cmd/server/cmd"
)

func main() {
	var cmd = cmd.NewRootCmd()
	err := cmd.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
