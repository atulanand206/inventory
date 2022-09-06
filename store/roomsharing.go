package store

import (
	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/types"
)

type roomSharingStore struct {
	MongoStore
}

type RoomSharingStore interface {
	CreateRooms(roomSharing []types.RoomSharing) error
}

func NewRoomSharingStore(config StoreConfig) RoomSharingStore {
	return &roomSharingStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *roomSharingStore) CreateRooms(roomSharing []types.RoomSharing) error {
	return m.Client.CreateMany(mapper.MapRoomSharingToInterface(roomSharing), m.Collection)
}
