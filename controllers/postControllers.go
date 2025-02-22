package controllers

import (
	"go-gin/initializers"
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(ctx *gin.Context) {
	// Validates user's input

	var body struct {
		Title string `json: "title" binding: "required" `
		Body string `json: "body" binding: required`
	}
	ctx.Bind(&body)

	// Get data off the request body creat a post and return it

	post:= models.Post{Title: body.Title, Body:body.Body}
	result:=initializers.DB.Create(&post)

	if(result.Error != nil ){
		ctx.JSON(http.StatusInternalServerError, gin.H{"message" : "Error creating post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})

}

func PostFindAll(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	ctx.JSON(http.StatusOK, gin.H{"message": "All posts", "posts": posts})
}