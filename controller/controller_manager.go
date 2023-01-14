package controller

import (
	"github.com/Hooneats/Syeong_server/controller/info"
	"github.com/Hooneats/Syeong_server/controller/user"
	"github.com/Hooneats/Syeong_server/service"
)

var InfoControl info.InfoController
var UserControl user.UserController

func InjectControllerDependency() {
	InfoControl = info.NewInfoControl()
	UserControl = user.NewUserControl(service.UserService)
}
