package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/kenkkd/ach-server/handler"
	"github.com/kenkkd/ach-server/pkg/config"
	"github.com/kenkkd/ach-server/pkg/ent"
)

func init() {
	config.Setup()
}

func main() {
	mysqlConfig := config.ApplicationConfig.MySQL
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	client, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	log.Println("success connecting to mysql")

	defer func(client *ent.Client) {
		err = client.Close()
		if err != nil {
			log.Fatalf("failed closing ent client: %v", err)
		}
	}(client)

	router := gin.Default()

	handler.HelloWorldHandler(router)

	err = router.Run()
	if err != nil {
		log.Fatalf("failed running router: %v", err)
	}
}
