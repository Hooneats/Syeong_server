package error

import (
	"errors"
)

type AppError struct {
	Code int
	MSG  string
	Name string
}

func (e AppError) Error() string {
	return e.MSG
}

func (e AppError) New() error {
	return errors.New(e.MSG)
}

func NewAppErrorAndMSG(MSG string) AppError {
	return NewAppErrorAndCode(UnKnownCode, MSG)
}

func NewAppError(e error) AppError {
	return NewAppErrorAndCode(UnKnownCode, e.Error())
}

func NewAppErrorAndCode(code int, MSG string) AppError {
	return NewAppErrorCustom(code, MSG, FailRequestErr)
}

func NewAppErrorCustom(code int, MSG string, name string) AppError {
	return AppError{
		Code: code,
		MSG:  MSG,
		Name: name,
	}
}

/* 401 ~ 499 error status */
var (
	UnauthorizedError         = NewAppErrorCustom(UnauthorizedCode, "허가되지 않은 사용자 입니다.", Unauthorized)
	RestAccessFailError       = NewAppErrorCustom(RestAccessFailCode, "로그인이 필요합니다.", RestAccessFail)
	BadRequestError           = NewAppErrorCustom(BadRequestCode, "부적절한 요청입니다.", BadRequest)
	DataNotFoundError         = NewAppErrorCustom(DataNotFoundCode, "데이터를 찾을 수 없습니다.", DataNotFound)
	DuplicatedDataError       = NewAppErrorCustom(DuplicatedDataCode, "이미 존재하는 데이터 입니다.", DuplicatedData)
	DoesNotModifyOrderError   = NewAppErrorCustom(DoesNotModifyOrderCode, "주문을 수정 할 수 없는 상태입니다.", DoesNotModifyOrder)
	BadAccessOrderError       = NewAppErrorCustom(DoesNotModifyOrderCode, "본인과 관련된 주문만 수정 가능 합니다.", DoesNotModifyOrder)
	AddNotRecommendMenusError = NewAppErrorCustom(AddNotRecommendMenusCode, "추가하려는 추천 메뉴들을 확인해 주세요.", AddNotRecommendMenus)
	DoesNotExistsOrderErr     = NewAppErrorCustom(DoesNotExistsOrderCode, "주문 기록이 없습니다.", DoesNotExistsOrder)
	DoesNotPostReviewErr      = NewAppErrorCustom(DoesNotPostReviewCode, "리뷰를 작성 할 수 없습니다.", DoesNotPostReview)
	FailDeleteRecommendErr    = NewAppErrorCustom(FailDeleteRecommendCode, "추천 메뉴에서 삭제하지 못 했습니다.", FailDeleteRecommend)
	UserNotFoundErr           = NewAppErrorCustom(UserNotFoundCode, "존재하지 않는 사용자입니다.", UserNotFound)
)

/* 501 ~ 599 서버 에러 */
var (
	NonInjectedError    = NewAppErrorCustom(InternalServerErrCode, "의존성 주입이 이루어지지 않았습니다.", NonInjected)
	InternalServerError = NewAppErrorCustom(SystemErrCode, "서버 로직을 수행하지 못했습니다.", InternalServerErr)
	FailRequestError    = NewAppErrorCustom(UnKnownCode, "요청을 정상적으로 처리하지 못했습니다.", FailRequestErr)
	UnKnownError        = NewAppErrorCustom(NonInjectedCode, "알 수 없는 오류입니다.", UnKnown)
	DoesNotReOrderError = NewAppErrorCustom(DoesNotOrderCode, "주문 취소는 되었지만, 재주문 하지 못했습니다.", DoesNotOrderErr)
)
