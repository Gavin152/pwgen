package cmd

import (
	"fmt"
	"pwgen/internal/random"

	"github.com/spf13/cobra"
)

var numberCount int
var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Generate a random number",
	Run: func(cmd *cobra.Command, args []string) {
		res := random.RandomIntAsString(numberCount)
		fmt.Println(res)
	},
}

func init() {
	numberCmd.Flags().IntVarP(&numberCount, "count", "c", 4, "Number of digits")
	RootCmd.AddCommand(numberCmd)
}
