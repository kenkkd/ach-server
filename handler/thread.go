package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kenkkd/ach-server/pkg/ent/thread"
	"github.com/kenkkd/ach-server/pkg/mysql"
)

func ThreadHandler(router *gin.Engine) {
	g := router.Group("/thread")
	g.GET("", getThreads)
}

func getThreads(c *gin.Context) {
	ctx := c.Request.Context()
	client := mysql.GetClient()

	res, err := client.Thread.Query().Where(thread.DeletedAtIsNil()).All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}
