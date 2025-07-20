package main

import (
	"booklibraryapi/config"
	_ "booklibraryapi/docs"
	"booklibraryapi/routes"
	"log"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book Library API
// @version 1.0
// @description REST API for managing books and categories
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token JWT dengan format: Bearer <token>
func main() {
	config.InitDB()
	runMigrations()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

func runMigrations() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(config.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Printf("Migration success. %d migration(s) applied.", n)
}
