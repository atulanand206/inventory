package store

import (
	"fmt"

	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type machineStore struct {
	mongoStore
	Collection string
}

type MachineStore interface {
	CreateMachines(machines []types.Machine)
	GetMachines() ([]types.Machine, error)
	MarkMachine(machine types.Machine) ([]types.Machine, error)
}

func NewMachineStoreConn(config StoreConfig) MachineStore {
	return &machineStore{
		mongoStore: mongoStore{
			client: Data(config),
		},
		Collection: config.TableName,
	}
}

func (m *machineStore) CreateMachines(machines []types.Machine) {
	for _, machine := range machines {
		m.client.Create(machine, m.Collection)
	}
}

func (m *machineStore) GetMachines() ([]types.Machine, error) {
	cursor, err := m.client.Find(m.Collection, bson.M{}, &options.FindOptions{})
	fmt.Println(cursor)
	if err != nil {
		return nil, err
	}
	return m.decodeMachines(cursor)
}

func (m *machineStore) MarkMachine(machine types.Machine) ([]types.Machine, error) {
	_, err := m.client.Update(m.Collection, bson.M{"id": machine.No}, machine)
	if err != nil {
		return nil, err
	}
	return m.GetMachines()
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
