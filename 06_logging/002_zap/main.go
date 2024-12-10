package main

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	main05()
}

// -- 基本的な使用例 -- //
func main01() {
	// デフォルトの Logger を作成
	logger, _ := zap.NewProduction() // プロダクション用
	defer logger.Sync()              // ログをフラッシュしてリソースを解放

	// 構造化ログの記録
	logger.Info("これは情報レベルのログです",
		zap.String("key", "value"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", 2),
	)

	// エラーレベルのログ
	logger.Error("エラーが発生しました",
		zap.String("原因", "無効な入力"),
	)
}

// -- 開発向けの簡略ロガー (SugaredLogger) -- //
func main02() {
	// SugaredLogger を作成
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Logger を SugaredLogger に変換
	sugar := logger.Sugar()

	// フォーマットを使ったログ
	sugar.Infof("ユーザー %s が %d 回目の試行に成功しました", "Alice", 3)

	// その他の簡略ログ
	sugar.Infow("試行成功",
		"user", "Alice",
		"attempt", 3,
	)
}

// -- カスタム設定 -- //
func main03() {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	defer logger.Sync()

	logger.Info("アプリケーションが起動しました")
}

// -- フィールドの型 -- //
func main04() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("詳細情報",
		zap.String("name", "テスト"),
		zap.Int("age", 25),
		zap.Bool("active", true),
		zap.Float64("score", 85.5),
		zap.Duration("elapsed", time.Second*10),
	)
}

// --  -- //
// エラーを返す処理を模擬するための関数
func doSomething() error {
	return errors.New("何かしらのエラーが発生")
}

func someFunction() error {
	// Logger の作成（実際のアプリケーションではグローバルに保持することが多い）
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	err = doSomething()
	if err != nil {
		// エラー情報を構造化してログ出力
		logger.Error("処理に失敗しました",
			zap.Error(err),
			zap.String("function", "someFunction"),
		)
		return err
	}
	return nil
}
func main05() {
	if err := someFunction(); err != nil {
		// main 関数でのエラーハンドリング
		// この例では単純にプログラムを終了させていますが、
		// 実際のアプリケーションでは適切なエラーハンドリングを行ってください
		return
	}
}
