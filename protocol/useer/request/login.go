package request

type Login struct {
	// TODO Name 을 추후 Email 로
	NicName  string `json:"nic_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
