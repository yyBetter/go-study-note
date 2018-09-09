package middleware

import (
	"github.com/gin-gonic/gin"
	"go-study-note/go-restful-api/demo9/handler"
	"go-study-note/go-restful-api/demo9/pkg/errno"
	"go-study-note/go-restful-api/demo9/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
