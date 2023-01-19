package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// RegisterCMD register connection cmd
// @return *cobra.Command
func RegisterCMD() *cobra.Command {
	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "register connection",
		Long:  "register connection information",
		RunE: func(cmd *cobra.Command, _ []string) error {
			registerConnection()
			fmt.Println("register success")
			return nil
		},
	}
	return registerCmd
}

// registerConnection using prompt to register a connection
func registerConnection() {
	var (
		role string
		ip   string
		err  error
	)

	//conn := connection.Connection{}

	rolePrompt := promptui.Prompt{
		Label: "Type role",
		Validate: func(input string) error {
			if len(input) <= 0 {
				return errors.New("please type valid role")
			}
			return nil
		},
	}

	ipPrompt := promptui.Prompt{
		Label: "Type ip",
		Validate: func(input string) error {
			if len(input) <= 0 {
				return errors.New("please type valid ip")
			}
			return nil
		},
	}

	role, err = rolePrompt.Run()
	ip, err = ipPrompt.Run()
	if err != nil {
		fmt.Println("register fail")
	}

	fmt.Printf("role: %s, ip: %s\n", role, ip)
}
