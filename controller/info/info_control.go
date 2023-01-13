package info

import (
	"github.com/Hooneats/Syeong_server/config"
	"github.com/Hooneats/Syeong_server/protocol"
	"github.com/gin-gonic/gin"
)

var instance *infoControl

type infoControl struct {
}

func NewInfoControl() *infoControl {
	if instance != nil {
		return instance
	}
	instance = &infoControl{}
	return instance
}

func (i *infoControl) GetInformation(c *gin.Context) {
	protocol.SuccessData(config.ServerConfig).Response(c)
}
