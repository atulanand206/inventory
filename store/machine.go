package store

import (
	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type machineStore struct {
	MongoStore
}

type MachineStore interface {
	CreateMachines(machine []types.Machine) error
	GetMachines(buildingId string) ([]types.Machine, error)
	GetMachine(id string) (types.Machine, error)
	UpdateMachine(machine types.Machine) ([]types.Machine, error)
}

func NewMachineStore(config StoreConfig) MachineStore {
	return &machineStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (ms *machineStore) CreateMachines(machines []types.Machine) error {
	return ms.Client.CreateMany(mapper.MapMachinesToInterface(machines), ms.Collection)
}

func (ms *machineStore) GetMachines(buildingId string) ([]types.Machine, error) {
	cursor, err := ms.Client.Find(ms.Collection, bson.M{"buildingId": buildingId}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return ms.decodeMachines(cursor)
}

func (ms *machineStore) GetMachine(id string) (raw types.Machine, err error) {
	doc, err := ms.Client.FindOne(ms.Collection, bson.M{"id": id}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = ms.decodeMachine(doc)
	if err != nil {
		return
	}
	return
}

func (ms *machineStore) UpdateMachine(machine types.Machine) ([]types.Machine, error) {
	_, err := ms.Client.Update(ms.Collection, bson.M{"id": machine.No}, machine)
	if err != nil {
		return nil, err
	}
	return ms.GetMachines(machine.BuildingId)
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

func (m *machineStore) decodeMachine(doc bson.Raw) (scope types.Machine, err error) {
	err = bson.Unmarshal(doc, &scope)
	return
}
