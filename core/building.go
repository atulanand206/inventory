package core

import "github.com/atulanand206/inventory/store"

type buildingService struct {
	buildingStore *store.MongoStore
}

type BuildingService interface {
}

func NewBuildingService(buildingConfig store.StoreConfig) BuildingService {
	return &buildingService{
		buildingStore: store.NewStoreConn(buildingConfig),
	}
}