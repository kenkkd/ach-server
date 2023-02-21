package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kenkkd/ach-server/handler"
	"github.com/kenkkd/ach-server/pkg/ent"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/ach")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	log.Println("success connecting to mysql")

	defer client.Close()

	router := gin.Default()

	handler.HelloWorldHandler(router)

	router.Run()
}
