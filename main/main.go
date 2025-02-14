package main

import (
	"alle-task/db"
)

func main() {
	db.InitDB()
	db.SeedTasks(db.DB)
}
