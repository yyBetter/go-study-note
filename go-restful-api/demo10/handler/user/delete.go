package user

import (
	"github.com/gin-gonic/gin"
	"go-study-note/go-restful-api/demo7/handler"
	"go-study-note/go-restful-api/demo7/model"
	"go-study-note/go-restful-api/demo7/pkg/errno"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
