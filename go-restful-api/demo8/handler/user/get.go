package user

import (
	"github.com/gin-gonic/gin"
	"go-study-note/go-restful-api/demo7/handler"
	"go-study-note/go-restful-api/demo7/model"
	"go-study-note/go-restful-api/demo7/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
