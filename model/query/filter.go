package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDefaultIDFilter(obID primitive.ObjectID) bson.M {
	return bson.M{"_id": obID}
}
