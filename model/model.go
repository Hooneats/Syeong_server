package model

import (
	"github.com/Hooneats/Syeong_server/common"
	"github.com/Hooneats/Syeong_server/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCollections map[string]*mongo.Collection
var instance *model

// model Collection 은 Store , Customer , Receipt , Review , user
type model struct {
	client *mongo.Client
}

func NewModel(URI, DBName string, colNames []string) (*model, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &model{}
	if err := instance.Connect(URI); err != nil {
		logger.AppLog.Error(err.Error())
		return nil, err
	}

	if err := instance.checkClient(); err != nil {
		return nil, err
	}

	instance.loadMongoCollections(colNames, DBName)

	return instance, nil
}

func (m *model) Connect(uri string) error {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	opt := options.Client().SetMaxPoolSize(100).SetTimeout(common.DatabaseClientTimeOut)
	client, err := mongo.Connect(ctx, opt.ApplyURI(uri))
	if err != nil {
		logger.AppLog.Error(err.Error())
		return err
	}

	instance.client = client
	return nil
}

// CreateIndexes 인덱스 생성
func (m *model) CreateIndexes(colName string, unique bool, indexName ...string) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var indexModels []mongo.IndexModel
	for _, name := range indexName {
		IDXModel := mongo.IndexModel{
			Keys: bson.M{name: 1}, Options: options.Index().SetUnique(unique),
		}
		indexModels = append(indexModels, IDXModel)
	}

	if _, err := MongoCollections[colName].Indexes().CreateMany(ctx, indexModels); err != nil {
		logger.AppLog.Error(err.Error())
		return
	}
}

// CreateCompoundIndex 복합 인텍스 생성
func (m *model) CreateCompoundIndex(colName string, unique bool, indexName ...string) {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	var aggregationIDXs []bson.E
	for _, name := range indexName {
		aggregationIDXs = append(aggregationIDXs, bson.E{name, 1})
	}

	IDXModel := mongo.IndexModel{
		Keys:    aggregationIDXs,
		Options: options.Index().SetUnique(unique),
	}

	if _, err := MongoCollections[colName].Indexes().CreateOne(ctx, IDXModel); err != nil {
		logger.AppLog.Error(err.Error())
		return
	}
}

func (m *model) checkClient() error {
	ctx, cancel := common.NewContext(common.ModelContextTimeOut)
	defer cancel()

	if err := m.client.Ping(ctx, nil); err != nil {
		logger.AppLog.Error(err.Error())
		return err
	}
	return nil
}

// loadMongoCollections 이미 (들어간)로드된 collection 은 넣지 않음
func (m *model) loadMongoCollections(colNames []string, DBName string) {
	for _, name := range colNames {
		m.putCollection(name, DBName)
	}
}

// putCollection 이미 (들어간)로드된 collection 은 넣지 않음
func (m *model) putCollection(collection, dbName string) {
	if MongoCollections == nil {
		MongoCollections = make(map[string]*mongo.Collection)
	}

	if _, exists := MongoCollections[collection]; exists {
		return
	}

	db := m.client.Database(dbName)
	col := db.Collection(collection)
	MongoCollections[collection] = col
}
