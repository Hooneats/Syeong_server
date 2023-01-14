package request

import (
	"github.com/Hooneats/Syeong_server/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PutPasswordUser struct {
	UserID   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *PutPasswordUser) ToPutPasswordUser(ID string) (*entity.User, error) {
	OBJID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:       OBJID,
		Password: r.Password,
	}, nil
}
