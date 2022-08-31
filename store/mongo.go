package store

import "github.com/atulanand206/inventory/types"

type mongoStore struct {
}

func NewMongoStore() DataStore {
	return &mongoStore{}
}

func (m *mongoStore) GetMachines() []types.Machine {
	return []types.Machine{}
}

func (m *mongoStore) MarkMachine(machine types.Machine) []types.Machine {
	return []types.Machine{}
}
