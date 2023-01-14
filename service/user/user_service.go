package user

import (
	"github.com/Hooneats/Syeong_server/common/ciper"
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/Hooneats/Syeong_server/config"
	"github.com/Hooneats/Syeong_server/model/user"
	"github.com/Hooneats/Syeong_server/protocol/useer/request"
	"github.com/Hooneats/Syeong_server/protocol/useer/response"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type userService struct {
	userModel user.UserModeler
}

var instance *userService

func NewUserService(modeler user.UserModeler) *userService {
	if instance != nil {
		return instance
	}
	instance = &userService{
		userModel: modeler,
	}
	return instance
}

func (u *userService) RegisterUser(user *request.PostUser) (string, error) {
	postUser := user.ToPostUser()
	postUser.Password = u.hashPassword(user.Password)
	savedID, err := u.userModel.PostUser(postUser)
	if err != nil {
		return enum.BlankSTR, err
	}
	return savedID, err
}

func (u *userService) ModifyUser(ID string, usr *request.PutUser) (int, error) {
	updateUser, err := usr.ToPutUser(ID)
	if err != nil {
		return 0, err
	}

	updateCount, err := u.userModel.UpdateUser(updateUser)
	if err != nil {
		return 0, err
	}
	return int(updateCount), nil
}

func (u *userService) FindUser(ID string) (*response.ResponseUser, error) {
	foundUser, err := u.userModel.SelectUser(ID)
	if err != nil {
		return nil, err
	}
	resUser := response.FromUser(foundUser)
	return resUser, nil
}

func (u *userService) DeleteUser(ID string) (int, error) {
	deletedCount, err := u.userModel.DeleteUser(ID)
	if err != nil {
		return 0, err
	}
	return int(deletedCount), nil
}

func (u *userService) hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (u *userService) verifyPassword(userPassword string, hashPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword)); err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(user *request.Login) (*response.Token, error) {
	foundUser, err := u.userModel.SelectUserByNicName(user.NicName)
	if err != nil {
		return nil, err
	}

	if err := u.verifyPassword(user.Password, foundUser.Password); err != nil {
		return nil, err
	}

	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := ciper.CreateToken(foundUser.ID.Hex(), accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return &response.Token{
		AccessToken:  token.AccessToken.Token,
		RefreshToken: token.RefreshToken.Token,
	}, nil
}
