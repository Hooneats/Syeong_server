package validator

import (
	"github.com/Hooneats/Syeong_server/common/enum"
	error2 "github.com/Hooneats/Syeong_server/common/error"
	"strings"
)

// CheckBlank string trim 한 값이 "" 인 경우 BadRequestError
func CheckBlank(STR string) error {
	s := strings.Trim(STR, " ")
	if s == enum.BlankSTR {
		return error2.BadRequestError
	}
	return nil
}
