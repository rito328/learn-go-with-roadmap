package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func requestTimingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now() // リクエストの開始時間を記録

		// 次のハンドラやミドルウェアを実行
		err := next(c)

		// リクエストの終了時間を記録し、処理時間を計算
		duration := time.Since(start)
		log.Printf("Path: %s | Method: %s | Duration: %s", c.Path(), c.Request().Method, duration)

		return err
	}
}

func NewServer() *echo.Echo {
	e := echo.New()

	// カスタムエラーハンドラー
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			log.Printf("Error occurred: code=%d, message=%v", he.Code, he.Message)
			c.JSON(he.Code, map[string]string{"error": he.Message.(string)})
		} else {
			log.Printf("Unexpected error: %v", err)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
	}

	// 通常ルート
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// エラーを発生させるルート
	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request Example")
	})

	return e
}

func main() {
	e := NewServer()

	//e := echo.New()
	//
	//e.HTTPErrorHandler = func(err error, c echo.Context) {
	//	// エラーの内容をターミナルに出力
	//	if he, ok := err.(*echo.HTTPError); ok {
	//		log.Printf("Error occurred: code=%d, message=%v", he.Code, he.Message)
	//	} else {
	//		log.Printf("Unexpected error: %v", err)
	//	}
	//
	//	// エラーレスポンスの設定
	//	c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred"})
	//}
	//
	//// ルートの定義
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//
	//// エラーを強制発生させるルート
	//e.GET("/error", func(c echo.Context) error {
	//	return echo.NewHTTPError(http.StatusBadRequest, "Bad Request Example")
	//})
	//
	//// グローバルミドルウェアを登録
	//e.Use(requestTimingMiddleware)
	//
	//e.GET("/hello", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, World!")
	//})
	//
	//// 特定のルートだけにミドルウェアを適用する
	//e.GET("/special", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Special route!")
	//}, requestTimingMiddleware)
	//
	//// JSON を返すハンドラー
	//e.GET("/user", func(c echo.Context) error {
	//	user := User{ID: 1, Name: "John Doe"}
	//	return c.JSON(http.StatusOK, user)
	//})
	//
	//e.GET("/user/:id", func(c echo.Context) error {
	//	id := c.Param("id") // 動的パラメータを取得
	//	return c.JSON(http.StatusOK, map[string]string{
	//		"user_id": id,
	//	})
	//})
	//
	//e.GET("/search", func(c echo.Context) error {
	//	keyword := c.QueryParam("keyword") // クエリパラメータを取得
	//	return c.JSON(http.StatusOK, map[string]string{
	//		"search": keyword,
	//	})
	//})
	//
	//e.POST("/user", func(c echo.Context) error {
	//	var user User
	//	// JSON を構造体にバインド
	//	if err := c.Bind(&user); err != nil {
	//		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	//	}
	//	// 追加処理例: ID を動的に設定
	//	user.ID = 1
	//	return c.JSON(http.StatusOK, user)
	//})
	//
	//api := e.Group("/api")
	//
	//// グループにミドルウェアを適用
	//api.Use(requestTimingMiddleware)
	//
	//api.GET("/users", func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, map[string]string{"message": "User list"})
	//})
	//api.GET("/users/:id", func(c echo.Context) error {
	//	id := c.Param("id")
	//	return c.JSON(http.StatusOK, map[string]string{"message": "User details for " + id})
	//})
	//
	//e.POST("/login", func(c echo.Context) error {
	//	var req struct {
	//		Username string `json:"username"`
	//		Password string `json:"password"`
	//	}
	//
	//	// JSON データのバインディングとエラーチェック
	//	if err := c.Bind(&req); err != nil {
	//		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON payload")
	//	}
	//
	//	// 必須フィールドのチェック
	//	if req.Username == "" || req.Password == "" {
	//		return echo.NewHTTPError(http.StatusBadRequest, "Username and password are required")
	//	}
	//
	//	// 成功レスポンス
	//	return c.JSON(http.StatusOK, map[string]string{"status": "logged in"})
	//})

	e.Logger.Fatal(e.Start(":8080"))
}
