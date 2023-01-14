package user

import "github.com/Hooneats/Syeong_server/model/entity"

type UserModeler interface {
	PostUser(user *entity.User) (string, error)

	UpdateUser(user *entity.User) (int64, error)

	SelectUser(id string) (*entity.User, error)

	SelectUserByNicName(nicName string) (*entity.User, error)

	DeleteUser(id string) (int64, error)
}
