package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type buildingStore struct {
	MongoStore
}

type BuildingStore interface {
	Create(building types.Building) error
	GetBuildings() ([]types.Building, error)
}

func NewBuildingStore(config StoreConfig) BuildingStore {
	return &buildingStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *buildingStore) Create(building types.Building) error {
	return m.Client.Create(building, m.Collection)
}

func (m *buildingStore) GetBuildings() ([]types.Building, error) {
	cursor, err := m.Client.Find(m.Collection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return m.decodeBuildings(cursor)
}

func (m *buildingStore) decodeBuildings(cursor []bson.Raw) (scopes []types.Building, err error) {
	for _, doc := range cursor {
		var scope types.Building
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}
