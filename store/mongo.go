package store

import (
	mongo "github.com/atulanand206/go-mongo"
)

var localDatabase = mongo.NewMockDb()

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
		return localDatabase
	}
	return mongo.NewDb(config.DbName)
}

func StoreConfigs(dbName string, collections []string, local bool) map[string]StoreConfig {
	var storeConfigs = make(map[string]StoreConfig)
	for _, collection := range collections {
		storeConfigs[collection] = StoreConfig{
			DbName:    dbName,
			TableName: collection,
			Local:     local,
		}
	}
	return storeConfigs
}
