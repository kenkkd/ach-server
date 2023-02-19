package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenkkd/ach-server/handler"
)

func main() {
	router := gin.Default()
	handler.HelloWorldHandler(router)

	router.Run()
}
