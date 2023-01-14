package user

import (
	"github.com/Hooneats/Syeong_server/common/error"
	"github.com/Hooneats/Syeong_server/common/validator"
	"github.com/Hooneats/Syeong_server/protocol"
	"github.com/Hooneats/Syeong_server/protocol/useer/request"
	"github.com/Hooneats/Syeong_server/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

var instance *userControl

type userControl struct {
	userService user.UserServicer
}

func NewUserControl(svc user.UserServicer) *userControl {
	if instance != nil {
		return instance
	}
	instance = &userControl{
		userService: svc,
	}
	return instance
}

func (u *userControl) GetUser(c *gin.Context) {
	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}

	user, err := u.userService.FindUser(userID)
	if err != nil {
		protocol.Fail(error.DataNotFoundError).Response(c)
		return
	}

	protocol.SuccessData(user).Response(c)
}

func (u *userControl) PutUser(c *gin.Context) {
	reqU := &request.PutUser{}
	if err := c.ShouldBindJSON(reqU); err != nil {
		protocol.Fail(error.BadRequestError).Response(c)
		return
	}

	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}

	cnt, err := u.userService.ModifyUser(userID, reqU)
	if err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"updated_count": cnt,
	}).Response(c)
}

func (u *userControl) DeleteUser(c *gin.Context) {
	userID := c.Query("user-id")
	if err := validator.CheckBlank(userID); err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}

	cnt, err := u.userService.DeleteUser(userID)
	if err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(gin.H{
		"deleted_count": cnt,
	}).Response(c)
}

func (u *userControl) PostUser(c *gin.Context) {
	reqU := &request.PostUser{}
	if err := c.ShouldBindJSON(reqU); err != nil {
		protocol.Fail(error.BadRequestError).Response(c)
		return
	}

	savedID, err := u.userService.RegisterUser(reqU)
	if err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(http.StatusCreated, gin.H{"saved_id": savedID}).Response(c)
}

func (u *userControl) Login(c *gin.Context) {
	reqL := &request.Login{}
	if err := c.ShouldBindJSON(reqL); err != nil {
		protocol.Fail(error.BadRequestError).Response(c)
		return
	}

	tokens, err := u.userService.Login(reqL)
	if err != nil {
		protocol.Fail(error.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(tokens).Response(c)
}
