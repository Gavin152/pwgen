package cmd

import (
	"fmt"
	"pwgen/pkg/random"

	"github.com/spf13/cobra"
)

var count int8
var randomCmd = &cobra.Command{
	Use: "number",
	Short: "Generate a random number",
	Run: func(cmd *cobra.Command, args []string) {
		res := random.RandomInt(count)
		fmt.Println(res)
	},
}



func init() {
	randomCmd.Flags().Int8VarP(&count, "digits", "d", 4, "Number of digits")
	rootCmd.AddCommand(randomCmd)
}
