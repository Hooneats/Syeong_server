package user

import (
	"github.com/Hooneats/Syeong_server/common"
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/Hooneats/Syeong_server/model/entity"
	"github.com/Hooneats/Syeong_server/model/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var instance *userModel

type userModel struct {
	collection *mongo.Collection
}

func NewUserModel(col *mongo.Collection) *userModel {
	if instance != nil {
		return instance
	}
	instance = &userModel{
		collection: col,
	}
	return instance
}

func (u *userModel) PostUser(user *entity.User) (string, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return enum.BlankSTR, err
	}

	return user.ID.Hex(), nil
}

func (u *userModel) UpdateUser(user *entity.User) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	filter := query.GetDefaultIDFilter(user.ID)
	setField := user.NewSetDUser()
	updateRes, err := u.collection.UpdateOne(ctx, filter, setField)
	if err != nil {
		return 0, err
	}

	return updateRes.ModifiedCount, nil
}

func (u *userModel) SelectUser(id string) (*entity.User, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *entity.User
	filter := query.GetDefaultIDFilter(obID)
	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userModel) SelectUserByNicName(nicName string) (*entity.User, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var user *entity.User
	filter := bson.M{"nic_name": nicName}
	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userModel) DeleteUser(id string) (int64, error) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := query.GetDefaultIDFilter(obID)
	delRes, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return delRes.DeletedCount, nil
}
