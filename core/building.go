package core

import "github.com/atulanand206/inventory/store"

type buildingService struct {
	buildingStore *store.BuildingStore
}

type BuildingService interface {
}

func NewBuildingService(config store.StoreConfig) BuildingService {
	return &buildingService{
		buildingStore: store.NewBuildingStoreConn(config),
	}
}
