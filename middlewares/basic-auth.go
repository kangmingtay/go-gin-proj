package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	// can replace this with call to db to retrieve username / password
	return gin.BasicAuth(gin.Accounts{
		"username": "password",
	})
}
