package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type accessStore struct {
	MongoStore
}

type AccessStore interface {
	CreateAccessCode(bedAccess types.BedAccess) error
	GetAccess(bedId string) (types.BedAccess, error)
}

func NewAccessStore(config StoreConfig) AccessStore {
	return &accessStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *accessStore) CreateAccessCode(bedAccess types.BedAccess) error {
	return m.Client.Create(bedAccess, m.Collection)
}

func (m *accessStore) GetAccess(bedId string) (raw types.BedAccess, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"bedId": bedId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeBedAccess(doc)
	if err != nil {
		return
	}
	return
}

func (m *accessStore) decodeBedAccess(doc bson.Raw) (scope types.BedAccess, err error) {
	err = bson.Unmarshal(doc, &scope)
	return
}
