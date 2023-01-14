package user

import "github.com/gin-gonic/gin"

type UserController interface {
	GetUser(c *gin.Context)

	PutUser(c *gin.Context)

	DeleteUser(c *gin.Context)

	PostUser(c *gin.Context)

	Login(c *gin.Context)
}
