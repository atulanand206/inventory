package store

type MachineStore struct {
	mongoStore
	Collection string
}

func NewMachineStoreConn(config StoreConfig) *MachineStore {
	return &MachineStore{
		mongoStore: mongoStore{
			Client: Data(config),
		},
		Collection: config.TableName,
	}
}
