package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "挨拶と起こし機能を提供するCLIアプリケーション",
	Long: `このアプリケーションは、以下の機能を提供します：

- hello:  指定した言語で挨拶をします（日本語・英語・韓国語・スペイン語）
- thank:  指定した言語で感謝を伝えます（日本語・英語・韓国語・スペイン語）
- wake:   寝ている人を起こします（優しく・強めに）

各コマンドは --help オプションで詳細な使用方法を確認できます。`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("rootコマンドが実行されました")
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
