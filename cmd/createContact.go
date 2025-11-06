/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createContactCmd represents the createContact command
var createContactCmd = &cobra.Command{
	Use:   "createContact <name> <email>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			// print as error
			fmt.Printf("Bad request! usage : createContact <name> <email>\n")
			os.Exit(1)
		}
		name := args[0]
		email := args[1]
		fmt.Printf("Creating contact: %s with email: %s\n", name, email)
	},
}

func init() {
	rootCmd.AddCommand(createContactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createContactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createContactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
