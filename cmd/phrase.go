package cmd

import (
	"fmt"
	"pwgen/internal/words"

	"github.com/spf13/cobra"
)

var wordCount int
var wordSeparator string

var phraseCmd = &cobra.Command{
	Use:   "phrase",
	Short: "Generate a phrase of multiple words, leveraging the words API",
	Long: `This generates a phrase or list of words. This function does not
	use a local dictionary, but fetches the words from the words API from 
	random-word-api.herokuapp.com. Please take this into consideration when
	generating phrases to use as passwords.`,
	Run: func(cmd *cobra.Command, args []string) {
		phrase := words.GetWords(wordCount, wordSeparator)
		fmt.Println(phrase)
	},
}

func init() {
	phraseCmd.Flags().StringVarP(&wordSeparator, "separator", "s", "-", "Separator bewteen words")
	phraseCmd.Flags().IntVarP(&wordCount, "count", "c", 3, "Number of words to fetch")
	rootCmd.AddCommand(phraseCmd)
}
