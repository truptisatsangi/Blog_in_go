package main

import (
	"go-gin/controllers"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testtools"
	"github.com/stretchr/testify/assert"
)

func TestMainRouter(t *testing.T) {
	// Set up Gin in test mode (to avoid spinning up a real server)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostFindAll)
	r.GET("/posts/:id", controllers.PostFindOne)
	r.PATCH("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	// Test for POST /posts
	t.Run("Create Post", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/posts", nil)
		w := testtools.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Post created")
	})
	// Test for GET /posts
	t.Run("Get All Posts", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/posts", nil)
		w := testtools.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Posts found")
	})
	// Test for GET /posts/:id
	t.Run("Get One Post", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/posts/1", nil)
		w := testtools.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Post found")
	})
	// Test for PATCH /post/:id
	t.Run("Update Post", func(t *testing.T) {
		req, _ := http.NewRequest("PATCH", "/post/1", nil)
		w := testtools.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Post updated")
	})
	// Test for DELETE /post/:id
	t.Run("Delete Post", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/post/1", nil)
		w := testtools.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Post deleted")
	})
}
