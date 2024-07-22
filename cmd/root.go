/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	//"pwgen/pkg/words"

	"github.com/spf13/cobra"
)

var number int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "quickly create a random password",
	Long: `A command line tool to quickly create a random password.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
//		Run: func(cmd *cobra.Command, args []string) {
//			words.GetWords(number);
//		},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().IntVarP(&number, "number", "n", 1, "Number of words to fetch")
}


