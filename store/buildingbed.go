package store

import (
	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type buildingBedStore struct {
	MongoStore
}

type BuildingBedStore interface {
	CreateBeds(buildingBeds []types.BuildingBed) error
	GetBedsByBuildingId(buildingId string) ([]types.BuildingBed, error)
}

func NewBuildingBedStore(config StoreConfig) BuildingBedStore {
	return &buildingBedStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *buildingBedStore) CreateBeds(buildingBeds []types.BuildingBed) error {
	return m.Client.CreateMany(mapper.MapBuildingBedsToInterface(buildingBeds), m.Collection)
}

func (m *buildingBedStore) GetBedsByBuildingId(buildingId string) (buildingBeds []types.BuildingBed, err error) {
	docs, err := m.Client.Find(m.Collection, bson.M{"buildingId": buildingId}, &options.FindOptions{})
	if err != nil {
		return
	}
	buildingBeds, err = m.decodeBuildingBeds(docs)
	if err != nil {
		return
	}
	return
}

func (m *buildingBedStore) decodeBuildingBeds(docs []bson.Raw) (buildingBeds []types.BuildingBed, err error) {
	for _, doc := range docs {
		var buildingBed types.BuildingBed
		err = bson.Unmarshal(doc, &buildingBed)
		if err != nil {
			return
		}
		buildingBeds = append(buildingBeds, buildingBed)
	}
	return
}
