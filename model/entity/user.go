package entity

import (
	"github.com/Hooneats/Syeong_server/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	NicName     string             `bson:"nic_name,omitempty"`
	Password    string             `bson:"password,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	Role        string             `bson:"role,omitempty"`
	PrivateKey  string             `bson:"private_key,omitempty"`
	Item        *dom.Item          `bson:"item,omitempty"`
	BaseTime    *dom.BaseTime      `bson:"base_time,omitempty"`
}

func (u *User) NewSetDUser() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"name", u.Name},
				{"nic_name", u.NicName},
				{"phone_number", u.PhoneNumber},
				{"role", u.Role},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (u *User) NewSetDUserPassword() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"password", u.Password},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}

func (u *User) NewSetDUserBlackIron() bson.D {
	return bson.D{
		{"$set",
			bson.D{
				{"item.black_iron", u.Item.BlackIron},
				{"base_time.updated_at", time.Now()},
			},
		},
	}
}
