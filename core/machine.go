package core

import (
	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type machineService struct {
	machineStore store.MachineStore
	usageStore   store.UsageStore
}

type MachineService interface {
	CreateMachines(request types.CreateMachinesRequest)
	GetMachines() ([]types.Machine, error)
	GetMachine(id string) (types.Machine, error)
	MarkMachine(machine types.MarkMachineRequest) ([]types.Machine, error)
	UnMarkMachine(machine types.MarkMachineRequest) ([]types.Machine, error)
}

func NewMachineService(machineConfig, usageConfig store.StoreConfig) MachineService {
	return &machineService{
		machineStore: store.NewMachineStore(machineConfig),
		usageStore:   store.NewUsageStore(usageConfig),
	}
}

func (m *machineService) CreateMachines(request types.CreateMachinesRequest) {
	m.machineStore.CreateMachines(mapper.MapCreateMachineRequestToMachines(request))
}

func (m *machineService) GetMachines() ([]types.Machine, error) {
	return m.machineStore.GetMachines()
}

func (m *machineService) GetMachine(id string) (types.Machine, error) {
	return m.machineStore.GetMachine(id)
}

func (m *machineService) MarkMachine(req types.MarkMachineRequest) ([]types.Machine, error) {
	machine, err := m.GetMachine(req.MachineId)
	if err != nil {
		return nil, err
	}
	_, err = m.usageStore.GetByMachineId(req.MachineId)
	if err != nil {
		return nil, err
	}
	return m.machineStore.UpdateMachine(machine)
}

func (m *machineService) UnMarkMachine(req types.MarkMachineRequest) ([]types.Machine, error) {
	machine, err := m.GetMachine(req.MachineId)
	if err != nil {
		return nil, err
	}
	_, err = m.usageStore.GetByMachineId(req.MachineId)
	if err != nil {
		return nil, err
	}
	return m.machineStore.UpdateMachine(machine)
}
