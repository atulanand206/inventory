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
	GetByMachineIds(machineIds []string) ([]types.Usage, error)
	SaveUsage(usage types.Usage) error
	DeleteUsage(machineId string) error
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

func (m *usageStore) GetByMachineIds(machineIds []string) (raw []types.Usage, err error) {
	docs, err := m.Client.Find(m.Collection, bson.M{"machineId": bson.M{"$in": machineIds}}, &options.FindOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeUsages(docs)
	if err != nil {
		return
	}
	return
}

func (m *usageStore) decodeUsages(cursor []bson.Raw) (scopes []types.Usage, err error) {
	for _, doc := range cursor {
		var scope types.Usage
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}

func (m *usageStore) SaveUsage(usage types.Usage) error {
	_, err := m.Client.Update(m.Collection, bson.M{"machineId": usage.MachineId}, usage)
	return err
}

func (m *usageStore) DeleteUsage(machineId string) error {
	_, err := m.Client.Delete(m.Collection, bson.M{"machineId": machineId})
	return err
}
