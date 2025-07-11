package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pedprojectFinal/internal/database"
	"pedprojectFinal/internal/handlers"
	"pedprojectFinal/internal/tasksService"
	"pedprojectFinal/internal/web/tasks"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	repo := tasksService.NewTaskRepository(db)
	service := tasksService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
