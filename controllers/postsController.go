package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/twk-mn/go-json-crud-api/initializers"
	"github.com/twk-mn/go-json-crud-api/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	// If error, return 400
	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return the post if success
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the posts (in an array)
	var posts []models.Post
	initializers.DB.Find((&posts))

	// Respond with all posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//delete post
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}
