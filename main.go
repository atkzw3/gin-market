package main

import "github.com/gin-gonic/gin"

// https://gin-gonic.com/ja/docs/quickstart/

// air パッケージでホットリロード
// https://github.com/air-verse/air
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
