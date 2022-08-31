package store

type buildingStore struct {
	mongoStore
	Collection string
}

type BuildingStore interface {
}

func NewBuildingStoreConn(config StoreConfig) BuildingStore {
	return &buildingStore{
		mongoStore: mongoStore{
			client: Data(config),
		},
		Collection: config.TableName,
	}
}
