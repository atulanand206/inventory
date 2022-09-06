package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userStore struct {
	MongoStore
}

type UserStore interface {
	GetUser(id string) (types.User, error)
	GetUsers(ids []string) ([]types.User, error)
}

func NewUserStore(config StoreConfig) UserStore {
	return &userStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *userStore) GetUser(id string) (raw types.User, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"id": id}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeUser(doc)
	if err != nil {
		return
	}
	return
}

func (m *userStore) GetUsers(ids []string) (scopes []types.User, err error) {
	cursor, err := m.Client.Find(m.Collection, bson.M{"id": bson.M{"$in": ids}}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return m.decodeUsers(cursor)
}

func (m *userStore) decodeUsers(cursor []bson.Raw) (scopes []types.User, err error) {
	for _, doc := range cursor {
		var scope types.User
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}

func (m *userStore) decodeUser(doc bson.Raw) (scope types.User, err error) {
	err = bson.Unmarshal(doc, &scope)
	return
}
