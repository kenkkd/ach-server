package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/kenkkd/ach-server/handler"
	"github.com/kenkkd/ach-server/pkg/config"
	"github.com/kenkkd/ach-server/pkg/ent"
	"github.com/kenkkd/ach-server/pkg/mysql"
)

func init() {
	config.Setup()
}

func main() {
	var err error
	client := mysql.NewClient()
	defer func(client *ent.Client) {
		err = client.Close()
		if err != nil {
			log.Fatalf("failed closing ent client: %v", err)
		}
	}(client)

	router := gin.Default()

	handler.ThreadHandler(router)
	handler.ResponseHandler(router)
	handler.HelloWorldHandler(router)

	err = router.Run()
	if err != nil {
		log.Fatalf("failed running router: %v", err)
	}
}
