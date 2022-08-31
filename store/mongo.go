package store

import (
	mongo "github.com/atulanand206/go-mongo"
)

type StoreConfig struct {
	DbName    string
	TableName string
	Local     bool
}

type mongoStore struct {
	Client mongo.DBConn
}

func Data(config StoreConfig) mongo.DBConn {
	if config.Local {
		return mongo.NewMockDb()
	}
	return mongo.NewDb(config.DbName)
}
