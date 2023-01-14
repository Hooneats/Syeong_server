package model

import (
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/Hooneats/Syeong_server/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppModel Modeler

var UserModel user.UserModeler

func LoadMongoModel(URI, DBName string, colNames []string) error {
	m, err := NewModel(URI, DBName, colNames)
	if err != nil {
		return err
	}
	AppModel = m
	return nil
}

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {
	UserModel = user.NewUserModel(m[enum.UserCollectionName])
}

func CreateIndexesInModels() {
	AppModel.CreateIndexes(enum.UserCollectionName, true, "nic_name")
}
