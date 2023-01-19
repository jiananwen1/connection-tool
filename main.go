package main

import (
	"connection_tool/cmd"
	"fmt"
	"github.com/spf13/cobra"
)

// functions:

// 1. register remote information

// 2. display remote information, can interactive select one remote

// 3. after select one, try to connect the remote

// 4. update remote information

func main() {
	fmt.Println("welcome to connection tool")

	mainCmd := &cobra.Command{Use: "conn"}
	mainCmd.AddCommand(cmd.RegisterCMD())

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
