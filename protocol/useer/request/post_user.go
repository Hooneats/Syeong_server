package request

import (
	"github.com/Hooneats/Syeong_server/model/entity"
	"github.com/Hooneats/Syeong_server/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostUser struct {
	Name        string `json:"name" binding:"required,min=2,max=15"`
	NicName     string `json:"nic_name" binding:"required,min=2,max=15"`
	Password    string `json:"password" binding:"required,min=13"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Role        string `json:"role"  binding:"required,eq=user|eq=admin"`
	PrivateKey  string `json:"private_key"` // TODO REQUIRED
}

func (r *PostUser) ToPostUser() *entity.User {
	return &entity.User{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		NicName:     r.NicName,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
		Role:        r.Role,
		PrivateKey:  r.PrivateKey,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
