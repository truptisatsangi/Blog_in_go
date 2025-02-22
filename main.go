package main

import (
	"go-gin/controllers"
	"go-gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB();
}

func main() {
	r := gin.Default();
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostFindAll)
	r.Run();
}