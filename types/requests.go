package types

type NewBuildingRequest struct {
	Name  string      `json:"name"`
	Rooms map[int]int `json:"rooms"`
}

type NewAddUserRequest struct {
	BuildingId    string        `json:"buildingId"`
	UserId        string        `json:"userId"`
	RoomNo        int           `json:"roomNo"`
	SharingStatus SharingStatus `json:"sharingStatus"`
	BedId         string        `json:"bedId"`
}

type NewRemoveUserRequest struct {
	BuildingId string `json:"buildingId"`
	UserId     string `json:"userId"`
}

type MarkMachineRequest struct {
	MachineId string `json:"machineId"`
	UserId    string `json:"userId"`
	Status    Status `json:"status"`
}
