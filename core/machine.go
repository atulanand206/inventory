package core

import (
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type machineService struct {
	machineStore *store.MongoStore
}

type MachineService interface {
	CreateMachines(machines []types.Machine)
	GetMachines() ([]types.Machine, error)
	MarkMachine(machine types.Machine) ([]types.Machine, error)
}

func NewMachineService(machineConfig store.StoreConfig) MachineService {
	return &machineService{
		machineStore: store.NewStoreConn(machineConfig),
	}
}

func (m *machineService) CreateMachines(machines []types.Machine) {
	for _, machine := range machines {
		m.machineStore.Client.Create(machine, m.machineStore.Collection)
	}
}

func (m *machineService) GetMachines() ([]types.Machine, error) {
	cursor, err := m.machineStore.Client.Find(m.machineStore.Collection, bson.M{}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return m.decodeMachines(cursor)
}

func (m *machineService) MarkMachine(machine types.Machine) ([]types.Machine, error) {
	_, err := m.machineStore.Client.Update(m.machineStore.Collection, bson.M{"id": machine.No}, machine)
	if err != nil {
		return nil, err
	}
	return m.GetMachines()
}

func (m *machineService) decodeMachines(cursor []bson.Raw) (scopes []types.Machine, err error) {
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
