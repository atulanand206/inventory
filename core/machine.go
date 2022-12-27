package core

import (
	"errors"

	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type machineService struct {
	machineStore     store.MachineStore
	usageStore       store.UsageStore
	bedUserStore     store.BedUserStore
	buildingBedStore store.BuildingBedStore
}

type MachineService interface {
	CreateMachines(request types.CreateMachinesRequest)
	GetMachines(buildingId string) ([]types.MachineUsage, error)
	GetMachine(id string) (types.Machine, error)
	MarkMachine(machine types.MarkMachineRequest) error
	UnMarkMachine(machine types.MarkMachineRequest) error
}

func NewMachineService(machineConfig, usageConfig, bedUserConfig, buildingBedConfig store.StoreConfig) MachineService {
	return &machineService{
		machineStore:     store.NewMachineStore(machineConfig),
		usageStore:       store.NewUsageStore(usageConfig),
		bedUserStore:     store.NewBedUserStore(bedUserConfig),
		buildingBedStore: store.NewBuildingBedStore(buildingBedConfig),
	}
}

func (m *machineService) CreateMachines(request types.CreateMachinesRequest) {
	m.machineStore.CreateMachines(mapper.MapCreateMachineRequestToMachines(request))
}

func (m *machineService) GetMachines(buildingId string) ([]types.MachineUsage, error) {
	machines, err := m.machineStore.GetMachines(buildingId)
	if err != nil {
		return nil, err
	}
	machineIds := make([]string, 0)
	for _, machine := range machines {
		machineIds = append(machineIds, machine.No)
	}
	usages, err := m.usageStore.GetByMachineIds(machineIds)
	if err != nil {
		return nil, err
	}
	usageMap := make(map[string]types.Usage)
	for _, usage := range usages {
		usageMap[usage.MachineId] = usage
	}
	machineUsages := make([]types.MachineUsage, 0)
	for _, machine := range machines {
		machineUsage := types.MachineUsage{
			Id:     machine.No,
			Name:   machine.Name,
			BedId:  usageMap[machine.No].BedId,
			Status: usageMap[machine.No].Status,
		}
		machineUsages = append(machineUsages, machineUsage)
	}
	return machineUsages, nil
}

func (m *machineService) GetMachine(id string) (types.Machine, error) {
	return m.machineStore.GetMachine(id)
}

func (m *machineService) MarkMachine(req types.MarkMachineRequest) error {
	bedUser, err := m.bedUserStore.GetBedUserByUserId(req.UserId)
	if err != nil {
		return errors.New("user not found in any bed")
	}
	buildingBed, err := m.buildingBedStore.GetByBedId(bedUser.BedId)
	if err != nil {
		return errors.New("bed not found in any building")
	}
	if buildingBed.BuildingId != req.BuildingId {
		return errors.New("user not in requested building")
	}
	usage, err := m.usageStore.GetByMachineId(req.MachineId)
	if err == nil {
		if bedUser.BedId == usage.BedId {
			return errors.New("machine already in use by requested user")
		}
		return errors.New("machine already in use by another user")
	}
	usage = mapper.MapToUsage(req.MachineId, bedUser.BedId, types.Busy)
	return m.usageStore.Create(usage)
}

func (m *machineService) UnMarkMachine(req types.MarkMachineRequest) error {
	bedUser, err := m.bedUserStore.GetBedUserByUserId(req.UserId)
	if err != nil {
		return errors.New("user not found in any bed")
	}
	buildingBed, err := m.buildingBedStore.GetByBedId(bedUser.BedId)
	if err != nil {
		return errors.New("bed not found in any building")
	}
	if buildingBed.BuildingId != req.BuildingId {
		return errors.New("user not in requested building")
	}
	usage, err := m.usageStore.GetByMachineId(req.MachineId)
	if err != nil {
		return errors.New("machine not in use")
	}
	if bedUser.BedId != usage.BedId {
		return errors.New("machine in use by another user, can't unmark")
	}
	return m.usageStore.DeleteUsage(usage.MachineId)
}
