package store

type BuildingStore struct {
	mongoStore
	Collection string
}

func NewBuildingStoreConn(config StoreConfig) *BuildingStore {
	return &BuildingStore{
		mongoStore: mongoStore{
			Client: Data(config),
		},
		Collection: config.TableName,
	}
}
