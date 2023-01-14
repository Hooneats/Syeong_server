package v1

import (
	"github.com/Hooneats/Syeong_server/controller"
	"github.com/gin-gonic/gin"
)

// Users ("app/v1/users")
func Users(usersUrl *gin.RouterGroup) {
	usersUrl.POST("/login", controller.UserControl.Login)
	usersUrl.POST("/user", controller.UserControl.PostUser)
	usersUrl.GET("/user", controller.UserControl.GetUser)
	usersUrl.PUT("/user", controller.UserControl.PutUser)
	usersUrl.DELETE("/user", controller.UserControl.DeleteUser)
}
