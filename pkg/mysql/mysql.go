package mysql

import (
	"fmt"
	"log"

	"github.com/kenkkd/ach-server/pkg/config"
	"github.com/kenkkd/ach-server/pkg/ent"
)

var client *ent.Client

func GetClient() *ent.Client {
	return client
}

func NewClient() *ent.Client {
	mysqlConfig := config.ApplicationConfig.MySQL
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	c, err := ent.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	client = c
	log.Println("success connecting to mysql")

	return client
}
