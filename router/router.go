package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router *gin.Engine := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080")
}
