package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	main06()
}

// -- 基本的な使い方 -- //
func main01() {
	// 標準出力にログを出力
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// 情報ログ
	log.Info().Msg("アプリケーションが起動しました")

	// 構造化ログ
	log.Info().
		Str("user", "test_user").
		Int("attempt", 3).
		Msg("ログイン試行")

	// エラーログ
	log.Error().
		Err(nil).
		Str("reason", "無効なパスワード").
		Msg("ログイン失敗")
}

// -- JSON 形式のログ -- //
func main02() {
	// JSON ログの例
	log.Info().
		Str("event", "user_signup").
		Str("user", "test_user").
		Msg("ユーザー登録が完了しました")
}

// -- ログレベルの設定 -- //
func main03() {
	// デフォルトは Info レベル
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	// このログは出力されない（Info レベル）
	log.Info().Msg("このログは表示されません")

	// このログは出力される（Warn レベル以上）
	log.Warn().Msg("警告ログです")
	log.Error().Msg("エラーログです")
}

// -- コンテキストの活用 -- //
func main04() {
	logger := log.With().
		Str("service", "auth-service").
		Str("version", "1.0.0").
		Logger()

	logger.Info().Msg("サービスを起動しました")
}

// -- エラーのログ記録 -- //
func someFunction() error {
	return fmt.Errorf("エラーになりました")
}
func main05() {
	err := someFunction()
	if err != nil {
		log.Error().
			Err(err).
			Str("function", "someFunction").
			Msg("処理に失敗しました")
	}
}

// -- ファイルにログを保存 -- //
func main06() {
	// ファイルを開く
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal().Err(err).Msg("ログファイルを作成できませんでした")
	}
	defer file.Close()

	// ファイルにログを出力
	log.Logger = log.Output(file)

	log.Info().Msg("ファイルにログを出力しました")
}
