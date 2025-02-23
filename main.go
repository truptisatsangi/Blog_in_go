package main

import (
	"go-gin/controllers"
	"go-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostFindAll)
	r.GET("/posts/:id", controllers.PostFindOne)
	r.PATCH("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run()
}
