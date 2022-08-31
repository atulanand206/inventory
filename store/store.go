package store

import "github.com/atulanand206/inventory/types"

type DataStore interface {
	GetMachines() []types.Machine
	MarkMachine(machine types.Machine) []types.Machine
}

func New() DataStore {
	var machines = []types.Machine{
		{Name: "Striker", No: 1, Status: types.Free},
		{Name: "Maverick", No: 2, Status: types.Free},
		{Name: "Diamond", No: 3, Status: types.Free},
		{Name: "Highland", No: 4, Status: types.Free},
	}
	return &dataStore{
		machines: machines,
	}
}

type dataStore struct {
	machines []types.Machine
}

func (ds *dataStore) GetMachines() []types.Machine {
	return ds.machines
}

func (ds *dataStore) MarkMachine(machine types.Machine) []types.Machine {
	for i, m := range ds.machines {
		if m.No == machine.No {
			ds.machines[i].Status = machine.Status
		}
	}
	return ds.machines
}
