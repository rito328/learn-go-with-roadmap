package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func myMiddleware(c *gin.Context) {
	fmt.Println("Middleware before handler")
	c.Next() // 次の処理を呼び出す
	fmt.Println("Middleware after handler")
}

// ルーティング設定を関数として切り出す
func setupRouter() *gin.Engine {
	// Gin のデフォルトのルーターを作成
	r := gin.Default()

	// ミドルウェアを全体に適用
	//r.Use(myMiddleware)

	r.GET("/special", myMiddleware, func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This route uses middleware"})
	})

	// GETエンドポイントの定義
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id") // 動的パラメータを取得
		c.JSON(200, gin.H{
			"user_id": id,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword") // クエリパラメータを取得
		c.JSON(200, gin.H{
			"search": keyword,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	api := r.Group("/api")
	api.Use(myMiddleware)
	{
		api.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Users endpoint"})
		})
	}

	return r
}

func main() {
	r := setupRouter()

	// サーバーの起動（デフォルトは8080ポート）
	err := r.Run()
	if err != nil {
		return
	}
}
