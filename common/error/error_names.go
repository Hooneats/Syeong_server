package error

const (
	Unauthorized         = "인증 오류"
	RestAccessFail       = "로그인 필요"
	BadRequest           = "부적절한 요청 오류"
	DataNotFound         = "존재하지 않는 데이터"
	DuplicatedData       = "이미 존재하는 데이터"
	DoesNotModifyOrder   = "주문 수정 불가"
	AddNotRecommendMenus = "추천 메뉴 수정 불가"
	DoesNotExistsOrder   = "주문 기록 없음"
	DoesNotPostReview    = "리뷰 작성 불가"
	FailDeleteRecommend  = "추천 메뉴 삭제 실패"
	UserNotFound         = "존재하지 않는 사용자"

	InternalServerErr = "서버 로직 오류"
	FailRequestErr    = "시스템 오류"
	UnKnown           = "알 수 없는 오류"
	DoesNotOrderErr   = "주문 오류"

	NonInjected = "의존성 주입 오류"
)

const (
	UnauthorizedCode         = 441
	RestAccessFailCode       = 442
	BadRequestCode           = 443
	DataNotFoundCode         = 444
	DuplicatedDataCode       = 445
	DoesNotModifyOrderCode   = 446
	AddNotRecommendMenusCode = 447
	DoesNotExistsOrderCode   = 448
	DoesNotPostReviewCode    = 449
	FailDeleteRecommendCode  = 450
	UserNotFoundCode         = 451

	InternalServerErrCode = 551
	SystemErrCode         = 552
	UnKnownCode           = 553
	DoesNotOrderCode      = 554

	NonInjectedCode = 554
)
