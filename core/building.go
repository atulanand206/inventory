package core

import (
	"errors"
	"fmt"

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
	AddUser(request types.NewAddUserRequest) (types.BedUser, error)
	RemoveUser(request types.NewRemoveUserRequest) (types.BedUser, error)
	GetUsers(buildingId string) ([]types.User, error)
	GetBedUsers(buildingId string) ([]types.BedUser, error)
	GetBuildingLayout(buildingId string) ([]types.BuildingBed, error)
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

func (m *buildingService) AddUser(request types.NewAddUserRequest) (types.BedUser, error) {
	user, err := m.userStore.GetUser(request.UserId)
	if err != nil {
		return types.BedUser{}, err
	}
	bed, err := m.bedUserStore.GetBed(request.BedId)
	if err != nil {
		return types.BedUser{}, err
	}
	if bed.UserId != "" {
		return types.BedUser{}, errors.New("bed already occupied")
	}
	if bed.UserId == user.Id {
		return types.BedUser{}, errors.New("user already in bed")
	}
	bedUser := mapper.MapCreateBedUser(request)
	err = m.bedUserStore.CreateBedUser(bedUser)
	if err != nil {
		return bedUser, err
	}
	return bedUser, nil
}

func (m *buildingService) RemoveUser(request types.NewRemoveUserRequest) (types.BedUser, error) {
	bedUser, err := m.bedUserStore.GetBedUserByUserId(request.UserId)
	if err != nil {
		return types.BedUser{}, err
	}
	bedUser.UserId = ""
	err = m.bedUserStore.DeleteBedUser(bedUser)
	if err != nil {
		return types.BedUser{}, err
	}
	return bedUser, nil
}

func (m *buildingService) GetUsers(buildingId string) ([]types.User, error) {
	bedUsers, err := m.GetBedUsers(buildingId)
	if err != nil {
		return []types.User{}, err
	}
	fmt.Println(bedUsers)
	userIds := mapper.MapBedUsersToUserIds(bedUsers)
	users, err := m.userStore.GetUsers(userIds)
	fmt.Println(users)
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
	fmt.Println(buildingBeds)
	bedIds := mapper.MapBuildingBedsToBedIds(buildingBeds)
	fmt.Println(bedIds)
	bedUsers, err := m.bedUserStore.GetUsersByBedIds(bedIds)
	fmt.Println(bedUsers)
	if err != nil {
		return []types.BedUser{}, err
	}
	return bedUsers, nil
}

func (m *buildingService) GetBuildingLayout(buildingId string) ([]types.BuildingBed, error) {
	return m.buildingBedStore.GetBedsByBuildingId(buildingId)
}
