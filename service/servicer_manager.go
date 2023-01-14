package service

import (
	"github.com/Hooneats/Syeong_server/model"
	"github.com/Hooneats/Syeong_server/service/user"
)

var UserService user.UserServicer

func InjectServicesDependency() {
	UserService = user.NewUserService(model.UserModel)
}
