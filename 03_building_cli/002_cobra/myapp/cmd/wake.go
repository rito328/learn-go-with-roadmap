package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var wakeName string
var wakeTime string

func init() {
	wakeCmd := &cobra.Command{
		Use:     "wake",
		Aliases: []string{"w"},
		Short:   "寝ている人を起こします",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("'gentle' または 'strong' サブコマンドを指定してください。")
			fmt.Println("使用例:")
			fmt.Println("  wake gentle --name 田中 --time 07:30")
			fmt.Println("  wake strong --name 田中")
		},
	}

	gentleCmd := &cobra.Command{
		Use:     "gentle",
		Aliases: []string{"g"},
		Short:   "優しく起こします",
		Run:     runWakeGentle,
	}

	strongCmd := &cobra.Command{
		Use:     "strong",
		Aliases: []string{"s"},
		Short:   "強めに起こします",
		Run:     runWakeStrong,
	}

	gentleCmd.Flags().StringVarP(&wakeName, "name", "n", "you", "起こす人の名前")
	gentleCmd.Flags().StringVarP(&wakeTime, "time", "t", "", "現在時刻（デフォルトは現在時刻）")

	strongCmd.Flags().StringVarP(&wakeName, "name", "n", "you", "起こす人の名前")
	strongCmd.Flags().StringVarP(&wakeTime, "time", "t", "", "現在時刻（デフォルトは現在時刻）")

	wakeCmd.AddCommand(gentleCmd)
	wakeCmd.AddCommand(strongCmd)
	rootCmd.AddCommand(wakeCmd)
}

func getCurrentTime() time.Time {
	currentTime := time.Now()
	if wakeTime != "" {
		parsedTime, err := time.Parse("15:04", wakeTime)
		if err == nil {
			currentTime = parsedTime
		}
	}
	return currentTime
}

func runWakeGentle(cmd *cobra.Command, args []string) {
	currentTime := getCurrentTime()
	hour := currentTime.Hour()

	var message string
	switch {
	case hour < 5:
		message = fmt.Sprintf("（小声で）%sさん、まだ深夜ですが...必要があって起こしました。申し訳ありません。", wakeName)
	case hour < 12:
		message = fmt.Sprintf("おはようございます、%sさん。素敵な朝をお迎えください。朝食の用意ができています。", wakeName)
	default:
		message = fmt.Sprintf("%sさん、お昼を過ぎていますよ。ゆっくり起きましょう。", wakeName)
	}

	fmt.Println(message)
}

func runWakeStrong(cmd *cobra.Command, args []string) {
	currentTime := getCurrentTime()
	hour := currentTime.Hour()

	var message string
	switch {
	case hour < 5:
		message = fmt.Sprintf("%sさん！！緊急事態です！！直ちに起きてください！！", wakeName)
	case hour < 12:
		message = fmt.Sprintf("%sさん！もう朝ですよ！！遅刻しますよ！！急いで起きてください！！", wakeName)
	default:
		message = fmt.Sprintf("%sさん！！こんな時間まで寝てるんですか！？早く起きてください！！", wakeName)
	}

	fmt.Println(message)
}
