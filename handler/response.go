package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	eResponse "github.com/kenkkd/ach-server/pkg/ent/response"
	eThread "github.com/kenkkd/ach-server/pkg/ent/thread"
	"github.com/kenkkd/ach-server/pkg/mysql"
)

func ResponseHandler(router *gin.Engine) {
	g := router.Group("/response")
	g.GET("",getResponses)
	g.POST("", createResponse)
}

type CreateResponseParams struct {
	ThreadID string `json:"threadId"`
	Name string `json:"name"`
	Content string `json:"content"`
}

func getResponses(c *gin.Context) {
	ctx := c.Request.Context()
	client := mysql.GetClient()

	res, err := client.Response.Query().Where(eResponse.DeletedAtIsNil()).All(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, res)
}

func createResponse(c *gin.Context) {
	ctx := c.Request.Context()
	client := mysql.GetClient()
	p := new(CreateResponseParams)

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threadExist, err := client.Thread.Query().Where(
		eThread.IDEQ(p.ThreadID),
		eThread.DeletedAtIsNil(),
	).Exist(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !threadExist {
		c.JSON(http.StatusNotFound, gin.H{"error": "thread not found"})
		return
	}

	count, err := client.Response.Query().Where(eResponse.HasThreadWith(eThread.IDEQ(p.ThreadID))).Count(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newNumber := count + 1

	response, err := client.Response.Create().SetThreadID(p.ThreadID).SetName(p.Name).SetContent(p.Content).SetNumber(newNumber).Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
