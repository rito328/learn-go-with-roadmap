package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	main0101()
}

// --  -- //
func main0101() {
	// 基本的なログ出力
	log.Println("アプリケーションを開始します") // 2024/12/09 10:30:45 アプリケーションを開始します

	// フォーマット付きのログ出力
	count := 3
	log.Printf("現在の数値: %d", count) // 2024/12/09 10:30:45 現在の数値: 3
}

func main0102() {
	// プレフィックスの設定
	log.SetPrefix("【INFO】")
	log.Println("システムチェックを開始") // 【INFO】2024/12/09 10:30:45 システムチェックを開始

	// プレフィックスの変更
	log.SetPrefix("【ERROR】")
	log.Println("エラーが発生しました") // 【ERROR】2024/12/09 10:30:45 エラーが発生しました
}

func main0103() {
	// ログファイルを作成
	file, err := os.OpenFile("./app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("ログファイルを開けません:", err)
	}
	defer file.Close()

	// 出力先をファイルに設定
	log.SetOutput(file)

	log.Println("このメッセージはファイルに書き込まれます")
}

func main0104() {
	// 日時のみを表示
	log.SetFlags(log.Ldate)
	log.Println("日付のみ") // 2024/12/09 メッセージ

	// 時刻のみを表示
	log.SetFlags(log.Ltime)
	log.Println("時刻のみ") // 10:30:45 メッセージ

	// 日時とファイル名、行番号を表示
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("詳細情報付き") // 2024/12/09 10:30:45 main.go:15: メッセージ
}

func main0105() {
	// 情報ログ用のロガー
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// エラーログ用のロガー
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("これは情報ログです")
	// INFO: 2024/12/09 10:30:45 これは情報ログです

	errorLog.Println("これはエラーログです")
	// ERROR: 2024/12/09 10:30:45 main.go:14: これはエラーログです
}

func somethingBadHappens() bool {
	return false
}
func unexpectedCondition() bool {
	return true
}
func main0106() {
	// Fatal - ログを出力して os.Exit(1) を呼び出す
	if somethingBadHappens() {
		log.Fatal("致命的なエラーが発生しました")
	}

	// Panic - ログを出力してパニックを発生させる
	if unexpectedCondition() {
		log.Panic("予期せぬエラーが発生しました")
	}
}

func main0107() {
	// ログの設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// ファイル操作の例
	file, err := os.Open("存在しないファイル.txt")
	if err != nil {
		log.Printf("ファイルオープンエラー: %v", err)
		return
	}
	defer file.Close()

	// 処理続行...
	log.Println("ファイルの処理を開始します")
}

// --  -- //
func main0201() {
	// 基本的なログ出力
	slog.Info("アプリケーションを開始します")

	// 属性（キーと値のペア）を持つログ出力
	slog.Info("ユーザーがログイン",
		"username", "taro",
		"ip", "192.168.1.1",
		"loginTime", time.Now(),
	)

	// 異なるログレベルの使用
	slog.Debug("デバッグ情報です")
	slog.Info("通常の情報です")
	slog.Warn("警告です")
	slog.Error("エラーが発生しました")
}

func main0202() {
	// JSONハンドラーの設定
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)

	// JSON形式でログが出力されます
	slog.Info("サーバー起動",
		"port", 8080,
		"env", "production",
	)

	// テキストハンドラーの設定
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	logger = slog.New(textHandler)
	slog.SetDefault(logger)

	// テキスト形式でログが出力されます
	slog.Info("サーバー起動",
		"port", 8080,
		"env", "production",
	)
}

func main0203() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug, // ログレベルの設定
		AddSource: true,            // ソースコードの位置情報を追加
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Debug("デバッグ情報が表示されます")
	slog.Info("ファイルの処理を開始", "filename", "test.txt")
}

func main0204() {
	// 属性グループを使用したログ出力
	slog.Info("データベース接続",
		slog.Group("database",
			slog.String("host", "localhost"),
			slog.Int("port", 5432),
			slog.String("user", "admin"),
		),
		slog.Group("settings",
			slog.Bool("ssl", true),
			slog.Int("timeout", 30),
		),
	)
}

func doSomething() error {
	return errors.New("何か問題が発生しました")
}
func main0205() {
	err := doSomething()
	if err != nil {
		slog.Error("操作に失敗しました",
			"error", err,
			"operation", "doSomething",
			"retry", false,
		)
	}
}

func main0206() {
	// コンテキストを持つロガーの作成
	ctx := context.Background()
	logger := slog.With(
		"service", "user-api",
		"version", "1.0.0",
	)

	// このロガーを使用する全てのログに上記の属性が付加されます
	logger.InfoContext(ctx, "リクエストを受信",
		"method", "GET",
		"path", "/users",
	)
}

func main0207() {
	// 開発環境用の設定
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	if os.Getenv("ENV") == "production" {
		// 本番環境ではInfoレベル以上のみ出力
		opts.Level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// 異なるレベルでのログ出力
	slog.Debug("詳細なデバッグ情報") // 開発環境でのみ表示
	slog.Info("通常の処理情報")      // 常に表示
	slog.Warn("警告メッセージ")      // 常に表示
	slog.Error("エラー情報")         // 常に表示
}

func main0208() {
	// ログファイルを開く
	file, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		slog.Error("ログファイルを開けません", "error", err)
		return
	}
	defer file.Close()

	// ファイルへのJSON形式での出力を設定
	handler := slog.NewJSONHandler(file, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// アプリケーションのログ出力
	slog.Info("アプリケーション開始",
		"env", os.Getenv("ENV"),
		"pid", os.Getpid(),
	)
}
