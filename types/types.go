package types

//
type User struct {
	Id    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
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
	Name string `json:"name" bson:"name"`
	No   string `json:"id" bson:"id"`
}

//
type Usage struct {
	Id        string `json:"id" bson:"id"`
	MachineId string `json:"machineId" bson:"machineId"`
	BedId     string `json:"bedId" bson:"bedId"`
	Status    Status `json:"status" bson:"status"`
}
