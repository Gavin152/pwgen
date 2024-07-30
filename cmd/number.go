package cmd

import (
	"fmt"
	"pwgen/internal/random"

	"github.com/spf13/cobra"
)

var numberCount int8
var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Generate a random number",
	Run: func(cmd *cobra.Command, args []string) {
		res := random.RandomInt(numberCount)
		fmt.Println(res)
	},
}

func init() {
	numberCmd.Flags().Int8VarP(&numberCount, "count", "c", 4, "Number of digits")
	RootCmd.AddCommand(numberCmd)
}
