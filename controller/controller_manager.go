package controller

import "github.com/Hooneats/Syeong_server/controller/info"

var InfoControl info.InfoController

func InjectControllerDependency() {
	InfoControl = info.NewInfoControl()
}
