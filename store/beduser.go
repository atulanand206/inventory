package store

import (
	"github.com/atulanand206/inventory/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bedUserStore struct {
	MongoStore
}

type BedUserStore interface {
	CreateBedUser(bedUser types.BedUser) error
	GetBed(bedId string) (types.BedUser, error)
	GetBedUserByUserId(userId string) (types.BedUser, error)
	GetByBuildingId(buildingId string) ([]types.BedUser, error)
	GetUsersByBedIds(bedIds []string) ([]types.BedUser, error)
	DeleteBedUser(bedUser types.BedUser) error
}

func NewBedUserStore(config StoreConfig) BedUserStore {
	return &bedUserStore{
		MongoStore: *NewStoreConn(config),
	}
}

func (m *bedUserStore) CreateBedUser(bedUser types.BedUser) error {
	return m.Client.Create(bedUser, m.Collection)
}

func (m *bedUserStore) GetBed(bedId string) (raw types.BedUser, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"bedId": bedId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeBedUser(doc)
	if err != nil {
		return
	}
	return
}

func (m *bedUserStore) GetBedUserByUserId(userId string) (raw types.BedUser, err error) {
	doc, err := m.Client.FindOne(m.Collection, bson.M{"userId": userId}, &options.FindOneOptions{})
	if err != nil {
		return
	}
	raw, err = m.decodeBedUser(doc)
	if err != nil {
		return
	}
	return
}

func (m *bedUserStore) GetByBuildingId(buildingId string) (bedUsers []types.BedUser, err error) {
	docs, err := m.Client.Find(m.Collection, bson.M{"buildingId": buildingId}, &options.FindOptions{})
	if err != nil {
		return
	}
	bedUsers, err = m.decodeBedUsers(docs)
	if err != nil {
		return
	}
	return
}

func (m *bedUserStore) GetUsersByBedIds(bedIds []string) (bedUsers []types.BedUser, err error) {
	docs, err := m.Client.Find(m.Collection, bson.M{"bedId": bson.M{"$in": bedIds}}, &options.FindOptions{})
	if err != nil {
		return
	}
	bedUsers, err = m.decodeBedUsers(docs)
	if err != nil {
		return
	}
	return
}

func (m *bedUserStore) DeleteBedUser(bedUser types.BedUser) error {
	_, err := m.Client.Delete(m.Collection, bson.M{"userId": bedUser.UserId})
	if err != nil {
		return err
	}
	return nil
}

func (m *bedUserStore) decodeBedUsers(cursor []bson.Raw) (scopes []types.BedUser, err error) {
	for _, doc := range cursor {
		var scope types.BedUser
		err = bson.Unmarshal(doc, &scope)
		if err != nil {
			return
		}
		scopes = append(scopes, scope)
	}
	return
}

func (m *bedUserStore) decodeBedUser(doc bson.Raw) (scope types.BedUser, err error) {
	err = bson.Unmarshal(doc, &scope)
	return
}
