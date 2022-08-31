package core

import (
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type buildingService struct {
	buildingStore store.BuildingStore
	userStore     store.UserStore
}

type BuildingService interface {
	Create(request types.NewBuildingRequest) (types.Building, error)
}

func NewBuildingService(buildingConfig store.StoreConfig, userConfig store.StoreConfig) BuildingService {
	return &buildingService{
		buildingStore: store.NewBuildingStore(buildingConfig),
		userStore:     store.NewUserStore(userConfig),
	}
}

func (m *buildingService) Create(request types.NewBuildingRequest) (types.Building, error) {
	var building types.Building
	building.Name = request.Name
	var id = ""
	building.Id = id
	err := m.buildingStore.Create(building)
	if err != nil {
		return building, err
	}
	return building, nil
}
