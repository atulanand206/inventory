package store

import (
	"github.com/atulanand206/inventory/role"
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userStore struct {
	MongoStore
}

type UserStore interface {
	CreateUser(user types.User) error
	GetUser(id string) (types.User, error)
	GetByUsername(username string) (types.User, error)
	GetUsers(ids []string) ([]types.User, error)
	GetByUsernames(usernames []string) ([]types.User, error)
	GetByRole(role role.Role) ([]types.User, error)
	UpdateUser(user types.User) error
}

func NewUserStore(config StoreConfig) UserStore {
	return &userStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *userStore) CreateUser(user types.User) error {
	return m.Client.Create(user, m.Collection)
}

func (m *userStore) GetUser(id string) (raw types.User, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"_id": id}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeUser(doc)
	if err != nil {
		return
	}
	return
}

func (m *userStore) GetByUsername(username string) (raw types.User, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"username": username}, &options.FindOneOptions{})
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
	cursor, err := m.Client.Find(m.Collection, bson.M{"_id": bson.M{"$in": ids}}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return m.decodeUsers(cursor)
}

func (m *userStore) GetByUsernames(usernames []string) (scopes []types.User, err error) {
	cursor, err := m.Client.Find(m.Collection, bson.M{"username": bson.M{"$in": usernames}}, &options.FindOptions{})
	if err != nil {
		return nil, err
	}
	return m.decodeUsers(cursor)
}

func (m *userStore) GetByRole(role role.Role) (scopes []types.User, err error) {
	cursor, err := m.Client.Find(m.Collection, bson.M{"role": role}, &options.FindOptions{})
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

func (m *userStore) UpdateUser(user types.User) error {
	_, err := m.Client.Update(m.Collection, bson.M{"_id": user.Id}, user)
	if err != nil {
		return err
	}
	return nil
}
