package main

import (
	"log"
	"os"
	"todolist/db"
	"todolist/internal/app/handler"
	"todolist/internal/app/repository"
	"todolist/internal/app/server"
	"todolist/internal/app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	dbUrl := os.Getenv("DB_URL")
	dbConn, err := db.New(dbUrl)
	if err != nil {
		log.Fatal("Error creating database connection:", err)
	}
	defer dbConn.Close()

	repository := repository.New(dbConn.GetDb())
	service := service.New(repository)
	handler := handler.New(service)
	server := server.New()

	server.UseMiddleware(gin.Logger())

	server.SetupRoutes(*handler)

	server.Run("0.0.0.0", "8000")
}
