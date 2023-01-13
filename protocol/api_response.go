package protocol

import (
	"fmt"
	error2 "github.com/Hooneats/Syeong_server/common/error"
	"github.com/Hooneats/Syeong_server/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse[T any] struct {
	Code      int    `json:"code"`
	Data      T      `json:"data"`
	Message   string `json:"message"`
	ErrorName string `json:"error"`
}

func (a *ApiResponse[T]) Response(c *gin.Context) {
	c.JSON(a.Code, a)
}
func Success() *ApiResponse[any] {
	return SuccessDataAndCustomMessage("success", "ok")
}

func SuccessData[T any](d T) *ApiResponse[any] {
	return SuccessDataAndCustomMessage(d, "ok")
}

func SuccessCodeAndData[T any](c int, d T) *ApiResponse[any] {
	return SuccessCustom(c, d, "ok")
}

func SuccessDataAndCustomMessage[T any](d T, MSG string) *ApiResponse[any] {
	return SuccessCustom(http.StatusOK, d, MSG)
}

func SuccessCustom[T any](c int, d T, MSG string) *ApiResponse[any] {
	return &ApiResponse[any]{
		Code:    c,
		Data:    d,
		Message: MSG,
	}
}

func Fail(e error2.AppError) *ApiResponse[any] {
	return FailCustomMessage(e, e.MSG)
}

func FailCustomMessage(e error2.AppError, MSG string) *ApiResponse[any] {
	errLogMSG := fmt.Sprintf("Error is %s, Code %d, Message : %s By %s", e.Name, e.Code, MSG, e.MSG)
	logger.AppLog.Error(errLogMSG)
	return &ApiResponse[any]{
		Code:      e.Code,
		Data:      nil,
		Message:   MSG,
		ErrorName: e.Name,
	}
}
