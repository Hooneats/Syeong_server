package model

import (
	// "fmt"
	// "encoding/json"
	"context"
	
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolModel struct {
	Client *mongo.Client
	PoolCollection *mongo.Collection
}

type Pool struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	City string
	Region string
	Name string
	Url string
	Address string
	Pnum string
	ImgUrl string
	LaneLength int
	LaneNum int
	CostInfoUrl string
	FreeSwimInfoUrl string
}

func GetPoolModel(db, host, model string) (*PoolModel, error) {
	pm := &PoolModel{}
	var err error

	if pm.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(host)); err != nil {
		return nil, err
	} else if err = pm.Client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	} else {
		pm.PoolCollection = pm.Client.Database(db).Collection(model)
	}

	return pm, nil
}