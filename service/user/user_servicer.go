package user

import (
	"github.com/Hooneats/Syeong_server/protocol/useer/request"
	"github.com/Hooneats/Syeong_server/protocol/useer/response"
)

type UserServicer interface {
	RegisterUser(user *request.PostUser) (string, error)

	ModifyUser(ID string, user *request.PutUser) (int, error)

	FindUser(ID string) (*response.ResponseUser, error)

	DeleteUser(ID string) (int, error)

	Login(user *request.Login) (*response.Token, error)
}
