package types

import "github.com/atulanand206/inventory/role"

//
type User struct {
	Id       string    `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	Phone    string    `json:"phone" bson:"phone"`
	Token    string    `json:"token" bson:"token"`
	Role     role.Role `json:"role" bson:"role"`
}

type CreateUserRequest struct {
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

type GetUserRequest struct {
	Username string `json:"username"`
}

type GetUsersRequest struct {
	Usernames []string `json:"usernames"`
}

type LoginRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type ResetPasswordRequest struct {
	Username    string `json:"username" bson:"username"`
	OldPassword string `json:"oldPassword" bson:"oldPassword"`
	NewPassword string `json:"newPassword" bson:"newPassword"`
}

type UserResponse struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
}

type BedAccess struct {
	BedId string `json:"bedId" bson:"bedId"`
	Code  string `json:"code" bson:"code"`
}

//
type Building struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type SharingStatus int

const (
	Single SharingStatus = iota
	Double
)

func (s SharingStatus) String() string {
	switch s {
	case Single:
		return "Single"
	case Double:
		return "Double"
	}
	return "unknown"
}

//
type RoomSharing struct {
	Id            string        `json:"id" bson:"id"`
	BuildingId    string        `json:"buildingId" bson:"buildingId"`
	RoomNo        int           `json:"roomNo" bson:"roomNo"`
	SharingStatus SharingStatus `json:"sharingStatus" bson:"sharingStatus"`
}

type OccStatus int

const (
	Occupied OccStatus = iota
	Vacant
)

func (s OccStatus) String() string {
	switch s {
	case Occupied:
		return "Occupied"
	case Vacant:
		return "Vacant"
	}
	return "unknown"
}

//
type BuildingBed struct {
	Id              string    `json:"id" bson:"id"`
	BuildingId      string    `json:"buildingId" bson:"buildingId"`
	BedId           string    `json:"bedId" bson:"bedId"`
	Floor           int       `json:"floor" bson:"floor"`
	RoomNo          int       `json:"roomNo" bson:"roomNo"`
	BedNo           int       `json:"bedNo" bson:"bedNo"`
	OccupancyStatus OccStatus `json:"occupancyStatus" bson:"occupancyStatus"`
}

type BuildingLayout struct {
	BuildingId string                                     `json:"buildingId" bson:"buildingId"`
	Layout     map[string]map[string]map[string]BedLayout `json:"layout" bson:"layout"`
}

type BedLayout struct {
	BedId  string `json:"bedId" bson:"bedId"`
	UserId string `json:"userId" bson:"userId"`
}

//
type BedUser struct {
	Id     string `json:"id" bson:"id"`
	BedId  string `json:"bedId" bson:"bedId"`
	UserId string `json:"userId" bson:"userId"`
}

type Status int

const (
	Free Status = iota
	Busy
	OutOfService
)

func (s Status) String() string {
	switch s {
	case Free:
		return "Free"
	case Busy:
		return "Busy"
	case OutOfService:
		return "OutOfService"
	}
	return "unknown"
}

//
type Machine struct {
	Name string `json:"name" bson:"name,unique"`
	No   string `json:"id" bson:"_id"`
}

//
type Usage struct {
	Id        string `json:"id" bson:"id"`
	MachineId string `json:"machineId" bson:"machineId"`
	BedId     string `json:"bedId" bson:"bedId"`
	Status    Status `json:"status" bson:"status"`
}

type MachineUsage struct {
	Id     string `json:"machineId" bson:"machineId"`
	Name   string `json:"name" bson:"name"`
	Status Status `json:"status" bson:"status"`
	BedId  string `json:"bedId" bson:"bedId"`
}
