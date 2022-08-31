package store

import "github.com/atulanand206/inventory/types"

type buildingStore struct {
	MongoStore
}

type BuildingStore interface {
	Create(building types.Building) error
}

func NewBuildingStore(config StoreConfig) BuildingStore {
	return &buildingStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *buildingStore) Create(building types.Building) error {
	return m.Client.Create(building, m.Collection)
}
