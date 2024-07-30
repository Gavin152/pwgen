/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	//"pwgen/internal/words"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "quickly create a random password",
	Long:  `A command line tool to quickly create a random password.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//		Run: func(cmd *cobra.Command, args []string) {
	//			words.GetWords(number);
	//		},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
