package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "挨拶します",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name", // name フラグ
						Value: "you",  // デフォルト値は "you"
						Usage: "相手の名前",
					},
					&cli.StringFlag{
						Name:  "lang",    // lang フラグ
						Value: "english", // デフォルト値は "english"
						Usage: "挨拶の言語",
					},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name") // Name フラグの値を取得

					var hello string
					lang := ctx.String("lang") // lang フラグの値を取得
					switch lang {
					case "japanese":
						hello = "こんにちは"
					case "korean":
						hello = "안녕하세요"
					case "spanish":
						hello = "Hola"
					default:
						hello = "Hello"
					}
					fmt.Printf("%s %s!!\n", hello, name)
					return nil
				},
			},
			{
				Name:  "thank",
				Usage: "感謝します",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name", // name フラグ
						Value: "you",  // デフォルト値は "you"
						Usage: "相手の名前",
					},
					&cli.StringFlag{
						Name:  "lang",    // lang フラグ
						Value: "english", // デフォルト値は "english"
						Usage: "挨拶の言語",
					},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name") // Name フラグの値を取得

					var thank string
					lang := ctx.String("lang") // lang フラグの値を取得
					switch lang {
					case "japanese":
						thank = "ありがとう"
					case "korean":
						thank = "고마워요"
					case "spanish":
						thank = "Gracias"
					default:
						thank = "Thank you"
					}
					fmt.Printf("%s %s!!\n", thank, name)
					return nil
				},
			},
			{
				Name:    "wake",
				Usage:   "寝ている人を起こします",
				Aliases: []string{"w"},
				Subcommands: []*cli.Command{
					{
						Name:    "gentle",
						Aliases: []string{"g"},
						Usage:   "優しく起こします",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Value:   "you",
								Usage:   "起こす人の名前",
							},
							&cli.StringFlag{
								Name:    "time",
								Aliases: []string{"t"},
								Value:   "",
								Usage:   "現在時刻（デフォルトは現在時刻）",
							},
						},
						Action: func(ctx *cli.Context) error {
							name := ctx.String("name")
							timeStr := ctx.String("time")

							currentTime := time.Now()
							if timeStr != "" {
								parsedTime, err := time.Parse("15:04", timeStr)
								if err == nil {
									currentTime = parsedTime
								}
							}

							var message string
							hour := currentTime.Hour()
							switch {
							case hour < 5:
								message = fmt.Sprintf("（小声で）%sさん、まだ深夜ですが...必要があって起こしました。申し訳ありません。", name)
							case hour < 12:
								message = fmt.Sprintf("おはようございます、%sさん。素敵な朝をお迎えください。朝食の用意ができています。", name)
							default:
								message = fmt.Sprintf("%sさん、お昼を過ぎていますよ。ゆっくり起きましょう。", name)
							}

							fmt.Println(message)
							return nil
						},
					},
					{
						Name:    "strong",
						Aliases: []string{"s"},
						Usage:   "強めに起こします",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Value:   "you",
								Usage:   "起こす人の名前",
							},
							&cli.StringFlag{
								Name:    "time",
								Aliases: []string{"t"},
								Value:   "",
								Usage:   "現在時刻（デフォルトは現在時刻）",
							},
						},
						Action: func(ctx *cli.Context) error {
							name := ctx.String("name")
							timeStr := ctx.String("time")

							currentTime := time.Now()
							if timeStr != "" {
								parsedTime, err := time.Parse("15:04", timeStr)
								if err == nil {
									currentTime = parsedTime
								}
							}

							var message string
							hour := currentTime.Hour()
							switch {
							case hour < 5:
								message = fmt.Sprintf("%sさん！！緊急事態です！！直ちに起きてください！！", name)
							case hour < 12:
								message = fmt.Sprintf("%sさん！もう朝ですよ！！遅刻しますよ！！急いで起きてください！！", name)
							default:
								message = fmt.Sprintf("%sさん！！こんな時間まで寝てるんですか！？早く起きてください！！", name)
							}

							fmt.Println(message)
							return nil
						},
					},
				},
				Action: func(ctx *cli.Context) error {
					fmt.Println("'gentle' または 'strong' サブコマンドを指定してください。")
					fmt.Println("使用例:")
					fmt.Println("  wake gentle --name 田中 --time 07:30")
					fmt.Println("  wake strong --name 田中")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func m01() {

}
