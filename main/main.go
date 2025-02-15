package main

import (
	"alle-task/db"
	"alle-task/handlers"
	"alle-task/middleware"
	"alle-task/repository"
	"alle-task/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Seed the database. Comment out if not needed.
	db.SeedTasks(db.DB)

	// Initialize the project repository, service modules and api handlers
	repo := repository.NewTaskRepository(db.DB)
	svc := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	// Initialize gin router
	r := gin.Default()

	// Initialize pagination middleware
	r.Use(middleware.Pagination())

	// API routes
	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetAllTasks)
	r.GET("/tasks/:id", handler.GetTaskByID)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	// Run the project
	r.Run(":8080")
}
