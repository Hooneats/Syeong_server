package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var AppModel Modeler

func LoadMongoModel(URI, DBName string, colNames []string) error {
	m, err := NewModel(URI, DBName, colNames)
	if err != nil {
		return err
	}
	AppModel = m
	return nil
}

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {

}

func CreateIndexesInModels() {

}
