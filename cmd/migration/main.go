package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kenkkd/ach-server/pkg/config"
	"github.com/kenkkd/ach-server/pkg/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.Setup()
}

func main() {
	mysqlConfig := config.ApplicationConfig.MySQL
	dataSourceName := fmt.Sprintf(
		"mysql://%s:%s@%s:%s/%s?charset=utf8&parseTime=True",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.MySQL),            // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
