package main

import (
	"fmt"
	"os"
	"xgmdr.com/pad/cmd/server/cmd"
)

func main() {
	var serverCommand = cmd.NewServerCmd()
	err := serverCommand.Execute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
