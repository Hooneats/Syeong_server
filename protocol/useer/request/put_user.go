package request

import (
	"github.com/Hooneats/Syeong_server/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PutUser struct {
	Name        string `json:"name" binding:"required,min=2,max=15"`
	NicName     string `json:"nic_name" binding:"required,min=2,max=15"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Role        string `json:"role"  binding:"required,eq=user|eq=admin"`
}

func (r *PutUser) ToPutUser(ID string) (*entity.User, error) {
	OBJID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:      OBJID,
		Name:    r.Name,
		NicName: r.NicName,
		Role:    r.Role,
	}, nil
}
