package role

type Role int

const (
	Admin Role = iota
	Manager
	Staff
	Customer
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Manager:
		return "Manager"
	case Staff:
		return "Staff"
	case Customer:
		return "Customer"
	}
	return "unknown"
}

func fromRoleString(role string) Role {
	switch role {
	case "Admin":
		return Admin
	case "Manager":
		return Manager
	case "Staff":
		return Staff
	case "Customer":
		return Customer
	}
	return -1
}

type Scope struct {
	Scopes []string `json:"scopes" bson:"scopes"`
}

type Capability int

const (
	User_Create Capability = iota
	User_Get
	Building_Create
	Building_Get_Overview
	Building_Get_Detailed
	Bed_Create
	Bed_Get
	Bed_Update
	BedUser_Create
	BedUser_Get
	BedUser_Update
	BedUser_Delete
	Machine_Create
	Machine_Get
	MachineUsage_Create
	MachineUsage_Delete
)

func (c Capability) String() string {
	switch c {
	case User_Create:
		return "user:create"
	case User_Get:
		return "user:get"
	case Building_Create:
		return "building:create"
	case Building_Get_Overview:
		return "building:get:overview"
	case Building_Get_Detailed:
		return "building:get:detailed"
	case Bed_Create:
		return "bed:create"
	case Bed_Get:
		return "bed:get"
	case Bed_Update:
		return "bed:update"
	case BedUser_Create:
		return "beduser:create"
	case BedUser_Get:
		return "beduser:get"
	case BedUser_Update:
		return "beduser:update"
	case BedUser_Delete:
		return "beduser:delete"
	case Machine_Create:
		return "machine:create"
	case Machine_Get:
		return "machine:get"
	case MachineUsage_Create:
		return "machineusage:create"
	case MachineUsage_Delete:
		return "machineusage:delete"
	}
	return "unknown"
}

func Capabilities(role Role) Scope {
	return Scope{
		Scopes: RoleCapabilities(role),
	}
}

func RoleCapabilities(role Role) []string {
	switch role {
	case Admin:
		return AdminCapabilities()
	case Manager:
		return ManagerCapabilities()
	case Staff:
		return StaffCapabilities()
	case Customer:
		return CustomerCapabilities()
	}
	return []string{}
}

func AdminCapabilities() []string {
	return []string{
		User_Create.String(),
		User_Get.String(),
		Building_Create.String(),
		Building_Get_Overview.String(),
		Building_Get_Detailed.String(),
		Bed_Create.String(),
		Bed_Get.String(),
		Bed_Update.String(),
		BedUser_Create.String(),
		BedUser_Get.String(),
		BedUser_Update.String(),
		BedUser_Delete.String(),
		Machine_Create.String(),
		Machine_Get.String(),
		MachineUsage_Create.String(),
		MachineUsage_Delete.String(),
	}
}

func ManagerCapabilities() []string {
	return []string{
		User_Create.String(),
		User_Get.String(),
		Building_Get_Overview.String(),
		Building_Get_Detailed.String(),
		Bed_Get.String(),
		Bed_Update.String(),
		BedUser_Create.String(),
		BedUser_Get.String(),
		BedUser_Update.String(),
		BedUser_Delete.String(),
		Machine_Get.String(),
	}
}

func StaffCapabilities() []string {
	return []string{
		User_Get.String(),
		Building_Get_Overview.String(),
		Building_Get_Detailed.String(),
		Bed_Get.String(),
		BedUser_Get.String(),
		Machine_Get.String(),
		MachineUsage_Delete.String(),
	}
}

func CustomerCapabilities() []string {
	return []string{
		User_Get.String(),
		Building_Get_Overview.String(),
		Building_Get_Detailed.String(),
		Bed_Get.String(),
		BedUser_Get.String(),
		Machine_Get.String(),
		MachineUsage_Create.String(),
		MachineUsage_Delete.String(),
	}
}

func HasCapability(roleString string, capability Capability) bool {
	role := fromRoleString(roleString)
	for _, c := range Capabilities(role).Scopes {
		if c == capability.String() {
			return true
		}
	}
	return false
}
