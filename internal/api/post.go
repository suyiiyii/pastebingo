package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pastebingo/dal/model"
	"pastebingo/dal/query"
	"strconv"
)

// PostCreate creates a new post.
//
//	@Summary		Create a new post
//	@Description	Create a new post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			post	body		model.Post	true	"Post"
//	@Success		200		{object}	model.Post
//	@Router			/posts [post]
func PostCreate(r *gin.RouterGroup) {
	r.POST("/posts", func(c *gin.Context) {
		var post model.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := query.Q.Post.Create(&post); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": post.ID})
	})
}

// PostGetList gets a post.
//
//	@Summary		Get a post
//	@Description	Get a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Post
//	@Router			/posts [get]
func PostGetList(r *gin.RouterGroup) {
	r.GET("/posts", func(c *gin.Context) {
		posts, err := query.Q.Post.Find()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})
}

// PostGet gets a post.
//
//	@Summary		Get a post
//	@Description	Get a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Post ID"
//	@Success		200		{object}	model.Post
//	@Router			/posts/{id} [get]
func PostGet(r *gin.RouterGroup) {
	r.GET("/posts/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		post, err := query.Q.Post.Where(query.Q.Post.ID.Eq(int32(id))).First()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})
}
