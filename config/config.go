// config/config.go
package config

import (
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func GetSqLDb() string {
	db := os.Getenv("SQL_DB")
	if db == "" {
		db = ""
	}
	return db
}
