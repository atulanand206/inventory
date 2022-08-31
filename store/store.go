package store

import "github.com/atulanand206/inventory/types"

type DataStore interface {
	GetMachines() []types.Machine
	MarkMachine(machine types.Machine) []types.Machine
}
