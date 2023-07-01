package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kenkkd/ach-server/pkg/ent/thread"
	"github.com/kenkkd/ach-server/pkg/mysql"
)

func ThreadHandler(router *gin.Engine) {
	g := router.Group("/thread")
	g.POST("", createThread)
	g.GET("", getThreads)
}

type CreateThreadParams struct {
	Title string `json:"title"`
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

func createThread(c *gin.Context) {
	ctx := c.Request.Context()
	client := mysql.GetClient()
	p := new(CreateThreadParams)

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := client.Thread.Create().SetTitle(p.Title).Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}
