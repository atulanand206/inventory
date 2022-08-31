package types

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

type Machine struct {
	Name   string `json:"name"`
	No     int    `json:"id"`
	Status Status `json:"status"`
}

type Inventory struct {
	Machines []Machine
}
