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
	db.InitDB()
	db.SeedTasks(db.DB)

	repo := repository.NewTaskRepository(db.DB)
	svc := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(svc)

	r := gin.Default()

	r.Use(middleware.Pagination())

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetAllTasks)
	r.GET("/tasks/:id", handler.GetTaskByID)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	r.Run(":8080")
}
