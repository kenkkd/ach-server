package handler

import "github.com/gin-gonic/gin"

func HelloWorldHandler(router *gin.Engine) {
	g := router.Group("/")
	g.GET("", helloWorld)
}

func helloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "Hello World!",
	})
}
