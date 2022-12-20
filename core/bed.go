package core

import (
	"errors"

	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type bedService struct {
	accessStore  store.AccessStore
	userStore    store.UserStore
	bedUserStore store.BedUserStore
}

type BedService interface {
	CreateBedAccess(bedAccess types.BedAccess) error
	ValidateAccess(bedAccess types.BedAccess) (bool, error)
	GetBedUser(bedId string) (types.BedUser, error)
	AddUser(request types.NewAddUserRequest) (types.BedUser, error)
	RemoveUser(request types.NewRemoveUserRequest) (types.BedUser, error)
}

func NewBedService(
	accessConfig,
	userConfig,
	bedUserConfig store.StoreConfig) BedService {
	return &bedService{
		accessStore:  store.NewAccessStore(accessConfig),
		userStore:    store.NewUserStore(userConfig),
		bedUserStore: store.NewBedUserStore(bedUserConfig),
	}
}

func (m *bedService) CreateBedAccess(bedAccess types.BedAccess) error {
	_, err := m.accessStore.GetAccess(bedAccess.BedId)
	if err == nil {
		return errors.New("code already exists")
	}
	bedAccess.Code = m.encryptAccessCode(bedAccess.Code)
	return m.accessStore.CreateAccessCode(bedAccess)
}

func (m *bedService) ValidateAccess(bedAccess types.BedAccess) (bool, error) {
	access, err := m.accessStore.GetAccess(bedAccess.BedId)
	if err != nil {
		return false, err
	}
	return access.Code == m.encryptAccessCode(bedAccess.Code), nil
}

func (m *bedService) encryptAccessCode(code string) string {
	return code
}

func (m *bedService) GetBedUser(bedId string) (types.BedUser, error) {
	return m.bedUserStore.GetBed(bedId)
}

func (m *bedService) AddUser(request types.NewAddUserRequest) (types.BedUser, error) {
	user, err := m.userStore.GetUser(request.UserId)
	if err != nil {
		return types.BedUser{}, err
	}
	bed, err := m.bedUserStore.GetBed(request.BedId)
	if err != nil {
		return types.BedUser{}, err
	}
	if bed.UserId != "" {
		return types.BedUser{}, errors.New("bed already occupied")
	}
	if bed.UserId == user.Id {
		return types.BedUser{}, errors.New("user already in bed")
	}
	bedUser := mapper.MapCreateBedUser(request)
	err = m.bedUserStore.CreateBedUser(bedUser)
	if err != nil {
		return bedUser, err
	}
	return bedUser, nil
}

func (m *bedService) RemoveUser(request types.NewRemoveUserRequest) (types.BedUser, error) {
	bedUser, err := m.bedUserStore.GetBedUserByUserId(request.UserId)
	if err != nil {
		return types.BedUser{}, err
	}
	bedUser.UserId = ""
	err = m.bedUserStore.DeleteBedUser(bedUser)
	if err != nil {
		return types.BedUser{}, err
	}
	return bedUser, nil
}
