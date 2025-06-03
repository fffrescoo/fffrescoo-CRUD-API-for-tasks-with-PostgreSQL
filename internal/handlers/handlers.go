package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"

	"pedprojectFinal/internal/models"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func GetHandler(c echo.Context) error {
	var tasks []models.Task
	if err := DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostHandler(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := DB.Create(task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, task)
}

func PatchHandler(c echo.Context) error {
	id := c.Param("id")
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
	}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}

func DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if err := DB.Delete(&models.Task{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
