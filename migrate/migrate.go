package main

import (
	"github.com/twk-mn/go-json-crud-api/initializers"
	"github.com/twk-mn/go-json-crud-api/models"
)

// This is the migration file
// Run: go run migrate/migrate.go

func init() {
	initializers.LoadEnvVariables() // Load secret
	initializers.ConnectToDB()      // Connect to the DB
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}) // "Seed" or Creates the DB structure
}
