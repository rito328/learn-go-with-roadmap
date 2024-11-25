package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var thankName string
var thankLang string

func init() {
	thankCmd := &cobra.Command{
		Use:   "thank",
		Short: "感謝します",
		Run:   runThank,
	}

	thankCmd.Flags().StringVarP(&thankName, "name", "n", "you", "相手の名前")
	thankCmd.Flags().StringVarP(&thankLang, "lang", "l", "english", "挨拶の言語")

	rootCmd.AddCommand(thankCmd)
}

func runThank(cmd *cobra.Command, args []string) {
	var thank string
	switch thankLang {
	case "japanese":
		thank = "ありがとう"
	case "korean":
		thank = "고마워요"
	case "spanish":
		thank = "Gracias"
	default:
		thank = "Thank you"
	}
	fmt.Printf("%s %s!!\n", thank, thankName)
}
