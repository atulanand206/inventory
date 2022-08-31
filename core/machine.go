package core

import (
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type machineService struct {
	machineStore store.MachineStore
}

type MachineService interface {
	CreateMachines(machines []types.Machine)
	GetMachines() ([]types.Machine, error)
	MarkMachine(machine types.Machine) ([]types.Machine, error)
}

func NewMachineService(machineConfig store.StoreConfig) MachineService {
	return &machineService{
		machineStore: store.NewMachineStore(machineConfig),
	}
}

func (m *machineService) CreateMachines(machines []types.Machine) {
	for _, machine := range machines {
		m.machineStore.Create(machine)
	}
}

func (m *machineService) GetMachines() ([]types.Machine, error) {
	return m.machineStore.GetMachines()
}

func (m *machineService) MarkMachine(machine types.Machine) ([]types.Machine, error) {
	return m.machineStore.UpdateMachine(machine)
}
