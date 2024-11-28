package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pastebingo/dal/model"
	"pastebingo/dal/query"
)

// UserCreate creates a new user.
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User	true	"User"
//	@Success		200		{object}	model.User
//	@Router			/users [post]
func UserCreate(r *gin.RouterGroup) {
	Q := query.Q
	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := Q.User.Create(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

}

// @Summary		Get all users
// @Description	Get all users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{array}		model.User
// @Router			/users [get]
func UserGet(r *gin.RouterGroup) {

	r.GET("/user", func(c *gin.Context) {
		users, err := query.Q.User.Find()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
}
