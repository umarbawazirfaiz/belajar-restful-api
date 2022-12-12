package main

import (
	"belajar-restful-api/database"
	"belajar-restful-api/exception"
	"belajar-restful-api/middleware"

	"belajar-restful-api/handler"
	"belajar-restful-api/repository"
	"belajar-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.GetConnection()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository, validate)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	echo := echo.New()
	api := echo.Group("/api")
	echo.Use(exception.PanicMiddleware, middleware.AuthMiddleware)
	api.GET("/categories", categoryHandler.FindAll)
	api.POST("/categories", categoryHandler.Create)
	api.PUT("/categories", categoryHandler.Update)
	api.DELETE("/categories", categoryHandler.Delete)
	api.GET("/categories/:id", categoryHandler.FindById)

	echo.Start(":3000")
}
