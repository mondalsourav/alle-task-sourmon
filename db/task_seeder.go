package db

import (
	"alle-task/models"
	"gorm.io/gorm"
	"log"
)

// SeedTasks adds dummy data to the tasks table
func SeedTasks(db *gorm.DB) {
	tasks := []models.Task{
		{Title: "Task 1", Description: "Description 1", Status: "Pending"},
		{Title: "Task 2", Description: "Description 2", Status: "In Progress"},
		{Title: "Task 3", Description: "Description 3", Status: "Completed"},
	}

	for _, task := range tasks {
		if err := db.Create(&task).Error; err != nil {
			log.Println("Failed to insert dummy task:", err)
		}
	}

	log.Println("Dummy data inserted successfully!")
}
