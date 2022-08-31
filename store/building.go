package store

type buildingStore struct {
	MongoStore
}

type BuildingStore interface {
}

func NewBuildingStore(config StoreConfig) BuildingStore {
	return &buildingStore{
		MongoStore: *NewStoreConn(config),
	}
}
