package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//func requestTimingMiddleware(c *fiber.Ctx) error {
//	start := time.Now() // リクエストの開始時間を記録
//
//	// 次のミドルウェアまたはハンドラーを実行
//	err := c.Next()
//
//	// リクエストの終了時間を記録し、処理時間を計算
//	duration := time.Since(start)
//	fmt.Printf("Path: %s | Method: %s | Duration: %s\n", c.Path(), c.Method(), duration)
//
//	return err
//}

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		// カスタムエラーハンドラー
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			// エラーレスポンスを記録
			fmt.Printf("Error occurred: code=%d, message=%s\n", code, err.Error())
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	// 通常ルート
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// エラーを発生させるルート
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request Example")
	})

	return app
}

func main() {
	// アプリケーションの初期化
	//app := fiber.New()
	//app := fiber.New(fiber.Config{
	//	// カスタムエラーハンドラーを設定
	//	ErrorHandler: func(c *fiber.Ctx, err error) error {
	//		// ステータスコードを取得 (デフォルトは 500)
	//		code := fiber.StatusInternalServerError
	//		if e, ok := err.(*fiber.Error); ok {
	//			code = e.Code
	//		}
	//
	//		// エラーの内容をログに出力
	//		fmt.Printf("Error occurred: code=%d, message=%v\n", code, err.Error())
	//
	//		// エラーレスポンスを返す
	//		return c.Status(code).JSON(fiber.Map{
	//			"error": "An error occurred",
	//		})
	//	},
	//})
	//
	//// 正常なルート
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//})
	//
	//// エラーを強制発生させるルート
	//app.Get("/error", func(c *fiber.Ctx) error {
	//	return fiber.NewError(fiber.StatusBadRequest, "Bad Request Example")
	//})
	//
	//// グローバルミドルウェアを登録
	////app.Use(requestTimingMiddleware)
	//
	//app.Get("/hello", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//})
	//
	//app.Get("/special", requestTimingMiddleware, func(c *fiber.Ctx) error {
	//	return c.SendString("Special route!")
	//})
	//
	//// ルートハンドラーの定義
	////app.Get("/", func(c *fiber.Ctx) error {
	////	return c.SendString("Hello, World!")
	////})
	//
	//app.Get("/user/:id", func(c *fiber.Ctx) error {
	//	id := c.Params("id") // 動的パラメータを取得
	//	return c.JSON(fiber.Map{
	//		"user_id": id,
	//	})
	//})
	//
	//app.Get("/search", func(c *fiber.Ctx) error {
	//	keyword := c.Query("keyword") // クエリパラメータを取得
	//	return c.JSON(fiber.Map{
	//		"search": keyword,
	//	})
	//})
	//
	//app.Post("/user", func(c *fiber.Ctx) error {
	//	// リクエストボディをパース
	//	var user struct {
	//		Name string `json:"name"`
	//		Age  int    `json:"age"`
	//	}
	//	if err := c.BodyParser(&user); err != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Invalid request body",
	//		})
	//	}
	//	return c.JSON(fiber.Map{
	//		"id":   1,
	//		"name": user.Name,
	//		"age":  user.Age,
	//	})
	//})
	//
	//api := app.Group("/api")
	//
	//// グループにミドルウェアを適用
	//api.Use(requestTimingMiddleware)
	//
	//api.Get("/users", func(c *fiber.Ctx) error {
	//	return c.JSON(fiber.Map{"message": "User list"})
	//})
	//api.Get("/users/:id", func(c *fiber.Ctx) error {
	//	id := c.Params("id")
	//	return c.JSON(fiber.Map{"message": "User details for " + id})
	//})
	//
	//app.Post("/login", func(c *fiber.Ctx) error {
	//	var req struct {
	//		Username string `json:"username"`
	//		Password string `json:"password"`
	//	}
	//
	//	// JSON データのパースとエラーチェック
	//	if err := c.BodyParser(&req); err != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Invalid JSON payload",
	//		})
	//	}
	//
	//	// 必須フィールドのチェック
	//	if req.Username == "" || req.Password == "" {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Username and password are required",
	//		})
	//	}
	//
	//	// 成功レスポンス
	//	return c.JSON(fiber.Map{
	//		"status": "logged in",
	//	})
	//})

	app := NewServer()

	// サーバーの起動
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
