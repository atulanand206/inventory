package store

type userStore struct {
	MongoStore
}

type UserStore interface {
}

func NewUserStore(config StoreConfig) UserStore {
	return &userStore{
		MongoStore: *NewStoreConn(config),
	}
}
