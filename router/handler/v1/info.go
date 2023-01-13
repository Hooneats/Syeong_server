package v1

import (
	"github.com/Hooneats/Syeong_server/controller"
	"github.com/gin-gonic/gin"
)

// Info ("app/v1/info")
func Info(infoUrl *gin.RouterGroup) {
	infoUrl.GET("", controller.InfoControl.GetInformation)
}
