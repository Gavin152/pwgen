package cmd

import (
	"fmt"
	"pwgen/pkg/random"

	"github.com/spf13/cobra"
)

var randomCount int
var randomLower, randomUpper, randomNumber, randomSymbol bool

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generate a random string",
	Run: func(cmd *cobra.Command, args []string) {
		options := random.RandomOptions{
			Lower:   randomLower,
			Upper:   randomUpper,
			Numbers: randomNumber,
			Symbols: randomSymbol,
		}
		if !options.Upper && !options.Numbers && !options.Symbols {
			options.Lower = true
		}
		res := random.RandomChars(randomCount, &options)
		fmt.Println(res)
	},
}

func init() {
	randomCmd.Flags().BoolVarP(&randomLower, "lowerCase", "l", false, "include lower case characters")
	randomCmd.Flags().BoolVarP(&randomUpper, "upperCase", "u", false, "include upper case characters")
	randomCmd.Flags().BoolVarP(&randomNumber, "numbers", "n", false, "include numbers")
	randomCmd.Flags().BoolVarP(&randomSymbol, "symbols", "s", false, "include special characters: ")
	randomCmd.Flags().IntVarP(&randomCount, "count", "c", 12, "Number of characters")
	rootCmd.AddCommand(randomCmd)
}
