package mapper

import (
	"github.com/atulanand206/inventory/types"
	"github.com/google/uuid"
)

func MapCreateBuilding(request types.NewBuildingRequest) types.Building {
	return types.Building{
		Id:   RandomUUId(),
		Name: request.Name,
	}
}

func MapCreateMachineRequestToMachines(request types.CreateMachinesRequest) []types.Machine {
	var machines []types.Machine
	for _, name := range request.Names {
		machines = append(machines, types.Machine{
			No:   RandomUUId(),
			Name: name,
		})
	}
	return machines
}

func MapMachinesToInterface(machine []types.Machine) []interface{} {
	var machinesInterface []interface{}
	for _, machine := range machine {
		machinesInterface = append(machinesInterface, machine)
	}
	return machinesInterface
}

func MapRoomSharingToInterface(roomSharing []types.RoomSharing) []interface{} {
	var roomSharingInterface []interface{}
	for _, room := range roomSharing {
		roomSharingInterface = append(roomSharingInterface, room)
	}
	return roomSharingInterface
}

func MapBuildingBedsToInterface(buildingBeds []types.BuildingBed) []interface{} {
	var buildingBedInterface []interface{}
	for _, bed := range buildingBeds {
		buildingBedInterface = append(buildingBedInterface, bed)
	}
	return buildingBedInterface
}

func MapCreateRoomShares(buildingId string, rooms map[int]int) []types.RoomSharing {
	var roomShares []types.RoomSharing
	for id, room := range rooms {
		for i := 1; i <= room; i++ {
			roomShares = append(roomShares, types.RoomSharing{
				Id:            RandomUUId(),
				RoomNo:        RoomNo(id, i),
				BuildingId:    buildingId,
				SharingStatus: types.Double,
			})
		}
	}
	return roomShares
}

func MapCreateBuildingBeds(roomShares []types.RoomSharing) []types.BuildingBed {
	var buildingBeds []types.BuildingBed
	for _, room := range roomShares {
		for i := 1; i <= 2; i++ {
			buildingBeds = append(buildingBeds, types.BuildingBed{
				Id:              RandomUUId(),
				BuildingId:      room.BuildingId,
				BedId:           RandomUUId(),
				Floor:           room.RoomNo / 100,
				RoomNo:          room.RoomNo,
				BedNo:           i,
				OccupancyStatus: types.Vacant,
			})
		}
	}
	return buildingBeds
}

func MapCreateBedUser(request types.NewAddUserRequest) types.BedUser {
	return types.BedUser{
		Id:     RandomUUId(),
		UserId: request.UserId,
		BedId:  request.BedId,
	}
}

func MapBuildingBedsToBedIds(buildingBeds []types.BuildingBed) []string {
	var userIds []string
	for _, bed := range buildingBeds {
		if bed.OccupancyStatus == types.Occupied {
			userIds = append(userIds, bed.BedId)
		}
	}
	return userIds
}

func MapBedUsersToUserIds(bedUsers []types.BedUser) []string {
	var userIds []string
	for _, user := range bedUsers {
		userIds = append(userIds, user.UserId)
	}
	return userIds
}

func RoomNo(floor int, no int) (roomNo int) {
	return floor*100 + no
}

func RandomUUId() string {
	return uuid.New().String()
}
