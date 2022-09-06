package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type usageStore struct {
	MongoStore
}

type UsageStore interface {
	Create(usage types.Usage) error
	GetByMachineId(machineId string) (types.Usage, error)
}

func NewUsageStore(config StoreConfig) UsageStore {
	return &usageStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *usageStore) Create(usage types.Usage) error {
	return m.Client.Create(usage, m.Collection)
}

func (m *usageStore) GetByMachineId(machineId string) (raw types.Usage, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"machineId": machineId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeUsage(doc)
	if err != nil {
		return
	}
	return
}

func (m *usageStore) decodeUsage(doc bson.Raw) (raw types.Usage, err error) {
	err = bson.Unmarshal(doc, &raw)
	return
}
