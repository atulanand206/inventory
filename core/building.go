package core

import "github.com/atulanand206/inventory/store"

type buildingService struct {
	buildingStore store.BuildingStore
	userStore     store.UserStore
}

type BuildingService interface {
}

func NewBuildingService(buildingConfig store.StoreConfig, userConfig store.StoreConfig) BuildingService {
	return &buildingService{
		buildingStore: store.NewBuildingStore(buildingConfig),
		userStore:     store.NewUserStore(userConfig),
	}
}
