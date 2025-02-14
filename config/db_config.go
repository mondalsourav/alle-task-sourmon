package config

import "alle-task/types"

var DevConfig = types.DBConfig{
	Host:     "localhost",
	Username: "postgres",
	Password: "admin123",
	Database: "alle-task-db",
}
