package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloName string
var helloLang string

func init() {
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "挨拶します",
		Run:   runHello,
	}

	helloCmd.Flags().StringVarP(&helloName, "name", "n", "you", "相手の名前")
	helloCmd.Flags().StringVarP(&helloLang, "lang", "l", "english", "挨拶の言語")

	rootCmd.AddCommand(helloCmd)
}

func runHello(cmd *cobra.Command, args []string) {
	var hello string
	switch helloLang {
	case "japanese":
		hello = "こんにちは"
	case "korean":
		hello = "안녕하세요"
	case "spanish":
		hello = "Hola"
	default:
		hello = "Hello"
	}
	fmt.Printf("%s %s!!\n", hello, helloName)
}
