package request

import (
	"github.com/Hooneats/Syeong_server/model/entity"
	"github.com/Hooneats/Syeong_server/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	UserID    string `json:"user_id"`
	BlackIron int    `json:"black_iron"`
}

func (i *Item) ToPutItemUser(ID string) (*entity.User, error) {
	OBJID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID: OBJID,
		Item: &dom.Item{
			BlackIron: i.BlackIron,
		},
	}, nil
}
