package main

import (
	"github.com/gin-gonic/gin"
	"github.com/twk-mn/go-json-crud-api/controllers"
	"github.com/twk-mn/go-json-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables() // Load secret
	initializers.ConnectToDB()      // Connect to the DB
}

func main() {
	r := gin.Default() // default gin router

	r.POST("/posts", controllers.PostsCreate)     // Create Post
	r.GET("/posts", controllers.PostIndex)        // Fetch all Posts
	r.GET("/post/:id", controllers.PostShow)      // Fetch by ID
	r.PUT("/post/:id", controllers.PostUpdate)    // Update Post (with ID)
	r.DELETE("/post/:id", controllers.PostDelete) // Delete Post (with ID)

	r.Run()
}
