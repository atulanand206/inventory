package mapper

import (
	"strconv"

	"github.com/atulanand206/inventory/role"
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

func MapCreateRoomShares(buildingId string, rooms map[string]int) []types.RoomSharing {
	var roomShares []types.RoomSharing
	for id, room := range rooms {
		flr, _ := strconv.Atoi(id)
		for i := 1; i <= room; i++ {
			roomShares = append(roomShares, types.RoomSharing{
				Id:            RandomUUId(),
				RoomNo:        RoomNo(flr, i),
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

func EncryptAccessCode(password string) string {
	return password
}

func MapCreateUser(request types.CreateUserRequest) types.User {
	return types.User{
		Id:       RandomUUId(),
		Name:     request.Name,
		Username: request.Username,
		Phone:    request.Phone,
		Role:     role.FromRoleString(request.Role),
		Token:    EncryptAccessCode(request.Password),
	}
}

func MapUserToResponse(user types.User) types.UserResponse {
	return types.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Phone:    user.Phone,
		Role:     user.Role,
	}
}

func MapCreateBedUser(request types.NewAddUserRequest) types.BedUser {
	return types.BedUser{
		Id:     RandomUUId(),
		UserId: request.UserId,
		BedId:  request.BedId,
	}
}

func MapBuildingBedsToBuildingLayout(buildingBeds []types.BuildingBed, userIds map[string]string) types.BuildingLayout {
	var buildingLayout types.BuildingLayout
	for _, bed := range buildingBeds {
		if buildingLayout.Layout == nil {
			buildingLayout.Layout = make(map[string]map[string]map[string]types.BedLayout)
		}
		if _, ok := buildingLayout.Layout[strconv.Itoa(bed.Floor)]; !ok {
			buildingLayout.Layout[strconv.Itoa(bed.Floor)] = make(map[string]map[string]types.BedLayout)
		}
		if _, ok := buildingLayout.Layout[strconv.Itoa(bed.Floor)][strconv.Itoa(bed.RoomNo)]; !ok {
			buildingLayout.Layout[strconv.Itoa(bed.Floor)][strconv.Itoa(bed.RoomNo)] = make(map[string]types.BedLayout)
		}
		var bedLayout types.BedLayout
		bedLayout.BedId = bed.BedId
		if ok := userIds[bed.BedId]; ok != "" {
			bedLayout.UserId = userIds[bed.BedId]
		}
		buildingLayout.Layout[strconv.Itoa(bed.Floor)][strconv.Itoa(bed.RoomNo)][strconv.Itoa(bed.BedNo)] = bedLayout
	}
	return buildingLayout
}

func MapBuildingBedsToBedIds(buildingBeds []types.BuildingBed) []string {
	var userIds []string
	for _, bed := range buildingBeds {
		// if bed.OccupancyStatus == types.Occupied {
		userIds = append(userIds, bed.BedId)
		// }
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
