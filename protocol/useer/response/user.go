package response

import (
	"github.com/Hooneats/Syeong_server/model/entity"
)

type ResponseUser struct {
	ID          string `json:"user_id"`
	Name        string `json:"name"`
	NicName     string `json:"nic_name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Item        Item   `json:"item"`
}

func FromUser(user *entity.User) *ResponseUser {
	return &ResponseUser{
		ID:          user.ID.Hex(),
		Name:        user.Name,
		NicName:     user.NicName,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		Item: Item{
			BlackIron: user.Item.BlackIron,
		},
	}
}
