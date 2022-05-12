package app

import "github.com/gin-gonic/gin"

func Startapplication() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}