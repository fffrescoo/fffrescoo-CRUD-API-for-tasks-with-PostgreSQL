package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"pedprojectFinal/internal/db"
	"pedprojectFinal/internal/handlers"
)

func main() {
	e := echo.New()

	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	handlers.SetDB(database)

	e.GET("/tasks", handlers.GetHandler)
	e.POST("/tasks", handlers.PostHandler)
	e.PATCH("/tasks/:id", handlers.PatchHandler)
	e.DELETE("/tasks/:id", handlers.DeleteHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
