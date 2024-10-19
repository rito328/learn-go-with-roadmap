package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	err := r.Run() // デフォルトで :8080 ポートで実行
	if err != nil {
		return
	}
}
