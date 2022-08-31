package store

import (
	mongo "github.com/atulanand206/go-mongo"
)

type StoreConfig struct {
	DbName    string
	TableName string
	Local     bool
}

type MongoStore struct {
	Client     mongo.DBConn
	Collection string
}

func NewStoreConn(config StoreConfig) *MongoStore {
	return &MongoStore{
		Client:     Data(config),
		Collection: config.TableName,
	}
}

func Data(config StoreConfig) mongo.DBConn {
	if config.Local {
		return mongo.NewMockDb()
	}
	return mongo.NewDb(config.DbName)
}
