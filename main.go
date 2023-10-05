package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ogibayashi/sample-app-golang/middleware"
)

func main() {
	// Ginルーターのインスタンスを作成
	r := gin.Default()
	r.Use(middleware.AccessLogMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	// GETエンドポイント
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// POSTエンドポイント
	r.POST("/post", func(c *gin.Context) {
		// リクエストからJSONデータをパース
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// レスポンスを返す
		c.JSON(http.StatusOK, gin.H{
			"message": "Received data",
			"data":    data,
		})
	})

	// サーバーを起動
	r.Run(":8080")
}
