package model

type Modeler interface {
	Connect(uri string) error
	CreateIndexes(colName string, unique bool, indexName ...string)
	CreateCompoundIndex(colName string, unique bool, indexName ...string)
}
