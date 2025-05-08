package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Task struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Task string    `json:"task"`
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("не удалось подключиться к базе данных: " + err.Error())
	}

	err = DB.AutoMigrate(&Task{})
	if err != nil {
		panic("не удалось провести миграцию: " + err.Error())
	}
}

func PostHandler(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if req.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task is empty"})
	}

	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	} else {
		var existingTask Task
		if result := DB.First(&existingTask, "id = ?", req.ID); result.Error == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID already exists"})
		}
	}

	if result := DB.Create(&req); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task"})
	}

	return c.JSON(http.StatusCreated, req)
}

func GetHandler(c echo.Context) error {
	var tasks []Task
	if result := DB.Find(&tasks); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func DeleteHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if result := DB.Delete(&Task{}, "id = ?", id); result.Error != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func PatchHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var update struct {
		Task string `json:"task"`
	}
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if update.Task == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task is empty"})
	}

	var task Task
	if result := DB.Model(&task).Where("id = ?", id).Update("task", update.Task); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	DB.First(&task, "id = ?", id)
	return c.JSON(http.StatusOK, task)
}

func main() {

	InitDB()

	e := echo.New()

	e.GET("/tasks", GetHandler)
	e.POST("/tasks", PostHandler)
	e.DELETE("/tasks/:id", DeleteHandler)
	e.PATCH("/tasks/:id", PatchHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
