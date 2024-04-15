package main

import (
	"database/sql"
	"example/nodejs-sqlite/drugs"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite Database
	db, err := sql.Open("sqlite3", "./ntsoekhe.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup drugs package and create table
	drugs.Setup(db)
	drugs.CreateTable()

	// Set up Gin
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS")
	r.Use(cors.New(corsConfig))

	// Routes
	r.GET("/drugs", drugs.ListDrugs)
	r.POST("/drugs", drugs.AddDrug)
	r.PUT("/drugs/:id", drugs.EditDrug)
	r.DELETE("/drugs/:id", drugs.DeleteDrug)

	// Start server
	r.Run() // listen and serve on 0.0.0.0:8080
}
