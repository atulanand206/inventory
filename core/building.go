package core

import (
	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type buildingService struct {
	buildingStore    store.BuildingStore
	roomSharingStore store.RoomSharingStore
	buildingBedStore store.BuildingBedStore
	bedUserStore     store.BedUserStore
	userStore        store.UserStore
}

type BuildingService interface {
	Create(request types.NewBuildingRequest) (types.Building, error)
	GetBuildings() ([]types.Building, error)
	GetUsers(buildingId string) ([]types.User, error)
	GetBedUsers(buildingId string) ([]types.BedUser, error)
	GetBuildingLayout(buildingId string) (types.BuildingLayout, error)
}

func NewBuildingService(
	bedUserConfig,
	buildingBedConfig,
	buildingConfig,
	roomSharingConfig,
	userConfig store.StoreConfig) BuildingService {
	return &buildingService{
		bedUserStore:     store.NewBedUserStore(bedUserConfig),
		buildingStore:    store.NewBuildingStore(buildingConfig),
		buildingBedStore: store.NewBuildingBedStore(buildingBedConfig),
		roomSharingStore: store.NewRoomSharingStore(roomSharingConfig),
		userStore:        store.NewUserStore(userConfig),
	}
}

func (m *buildingService) Create(request types.NewBuildingRequest) (types.Building, error) {
	building := mapper.MapCreateBuilding(request)
	err := m.buildingStore.Create(building)
	if err != nil {
		return building, err
	}
	roomShares := mapper.MapCreateRoomShares(building.Id, request.Rooms)
	err = m.roomSharingStore.CreateRooms(roomShares)
	if err != nil {
		return building, err
	}
	buildingBeds := mapper.MapCreateBuildingBeds(roomShares)
	err = m.buildingBedStore.CreateBeds(buildingBeds)
	if err != nil {
		return building, err
	}
	return building, nil
}

func (m *buildingService) GetBuildings() ([]types.Building, error) {
	return m.buildingStore.GetBuildings()
}

func (m *buildingService) GetUsers(buildingId string) ([]types.User, error) {
	bedUsers, err := m.GetBedUsers(buildingId)
	if err != nil {
		return []types.User{}, err
	}
	userIds := mapper.MapBedUsersToUserIds(bedUsers)
	users, err := m.userStore.GetUsers(userIds)
	if err != nil {
		return []types.User{}, err
	}
	return users, nil
}

func (m *buildingService) GetBedUsers(buildingId string) ([]types.BedUser, error) {
	buildingBeds, err := m.buildingBedStore.GetBedsByBuildingId(buildingId)
	if err != nil {
		return []types.BedUser{}, err
	}
	bedIds := mapper.MapBuildingBedsToBedIds(buildingBeds)
	bedUsers, err := m.bedUserStore.GetUsersByBedIds(bedIds)
	if err != nil {
		return []types.BedUser{}, err
	}
	return bedUsers, nil
}

func (m *buildingService) GetBuildingLayout(buildingId string) (types.BuildingLayout, error) {
	building, err := m.buildingStore.GetById(buildingId)
	if err != nil {
		return types.BuildingLayout{}, err
	}
	beds, err := m.buildingBedStore.GetBedsByBuildingId(buildingId)
	if err != nil {
		return types.BuildingLayout{}, err
	}
	bedIds := mapper.MapBuildingBedsToBedIds(beds)
	bedUsers, err := m.bedUserStore.GetUsersByBedIds(bedIds)
	if err != nil {
		return types.BuildingLayout{}, err
	}
	userIds := make(map[string]string)
	for _, bedUser := range bedUsers {
		userIds[bedUser.BedId] = bedUser.UserId
	}
	layout := mapper.MapBuildingBedsToBuildingLayout(building, beds, userIds)
	return layout, nil
}
