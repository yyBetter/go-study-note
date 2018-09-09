package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"go-study-note/go-restful-api/demo5/pkg/errno"
	"go-study-note/go-restful-api/demo6/hadler"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&r); err != nil {
		hadler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("username: %s", admin2)
	desc := c.Query("desc")
	log.Infof("param desc: %s", desc)
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		hadler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		hadler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	hadler.SendResponse(c, nil, rsp)
}
