package cmd

import (
	"fmt"
	"pwgen/pkg/random"

	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use: "random",
	Short: "Generate a random number",
	Run: func(cmd *cobra.Command, args []string) {
		res := random.RandomInt()
		fmt.Println(res)
	},
}

func init() {
	RootCmd.AddCommand(randomCmd)
}
