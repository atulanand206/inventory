package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type machineStore struct {
	MongoStore
}

type MachineStore interface {
	Create(machine types.Machine) error
	GetMachines() ([]types.Machine, error)
	UpdateMachine(machine types.Machine) ([]types.Machine, error)
}

func NewMachineStore(config StoreConfig) MachineStore {
	return &machineStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (ms *machineStore) Create(machine types.Machine) error {
	return ms.Client.Create(machine, ms.Collection)
}

func (ms *machineStore) GetMachines() ([]types.Machine, error) {
	cursor, err := ms.Client.Find(ms.Collection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return ms.decodeMachines(cursor)
}

func (ms *machineStore) UpdateMachine(machine types.Machine) ([]types.Machine, error) {
	_, err := ms.Client.Update(ms.Collection, bson.M{"id": machine.No}, machine)
	if err != nil {
		return nil, err
	}
	return ms.GetMachines()
}

func (m *machineStore) decodeMachines(cursor []bson.Raw) (scopes []types.Machine, err error) {
	for _, doc := range cursor {
		var scope types.Machine
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}
