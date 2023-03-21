package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenkkd/ach-server/pkg/ent/response"
	"github.com/kenkkd/ach-server/pkg/mysql"
)

func ResponseHandler(router *gin.Engine) {
	g := router.Group("/response")
	g.POST("", createResponse)
	g.GET("",getResponses)
}

func createResponse(c *gin.Context){
	ctx := c.Request.Context()
	client := mysql.GetClient()
	response,err := client.Response.Create().SetName("test").SetContent("test").SetThreadID(1).SetNumber(1).Save(ctx)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	c.JSON(http.StatusOK, response)
}

func getResponses(c *gin.Context) {
	ctx := c.Request.Context()
	client := mysql.GetClient()

	res, err := client.Response.Query().Where(response.DeletedAtIsNil()).All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}
