package core

import (
	"errors"

	"github.com/atulanand206/inventory/mapper"
	"github.com/atulanand206/inventory/role"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type userService struct {
	userStore store.UserStore
}

type UserService interface {
	CreateUser(user types.CreateUserRequest) (types.UserResponse, error)
	GetUser(username string) (types.UserResponse, error)
	GetUsers(usernames []string) ([]types.UserResponse, error)
	GetCustomers() ([]types.UserResponse, error)
	LoginUser(request types.LoginRequest) (types.UserResponse, error)
	ResetPassword(request types.ResetPasswordRequest) error
}

func NewUserService(userStore store.UserStore) UserService {
	return &userService{
		userStore: userStore,
	}
}

func (m *userService) CreateUser(userRequest types.CreateUserRequest) (types.UserResponse, error) {
	user := mapper.MapCreateUser(userRequest)
	m.userStore.CreateUser(user)
	return m.GetUser(user.Username)
}

func (m *userService) GetUser(username string) (types.UserResponse, error) {
	user, err := m.userStore.GetByUsername(username)
	if err != nil {
		return types.UserResponse{}, err
	}
	return mapper.MapUserToResponse(user), nil
}

func (m *userService) GetUsers(usernames []string) ([]types.UserResponse, error) {
	users, err := m.userStore.GetByUsernames(usernames)
	if err != nil {
		return nil, errors.New("users not available")
	}
	var userResponses []types.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, mapper.MapUserToResponse(user))
	}
	return userResponses, nil
}

func (m *userService) GetCustomers() ([]types.UserResponse, error) {
	users, err := m.userStore.GetByRole(role.Customer)
	if err != nil {
		return nil, errors.New("users not available")
	}
	var userResponses []types.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, mapper.MapUserToResponse(user))
	}
	return userResponses, nil
}

func (m *userService) LoginUser(request types.LoginRequest) (types.UserResponse, error) {
	token := mapper.EncryptAccessCode(request.Password)
	user, err := m.userStore.GetByUsername(request.Username)
	if err != nil {
		return types.UserResponse{}, errors.New("user not present")
	}
	if user.Token != token {
		return types.UserResponse{}, errors.New("invalid credentials")
	}
	return mapper.MapUserToResponse(user), nil
}

func (m *userService) ResetPassword(request types.ResetPasswordRequest) error {
	user, err := m.userStore.GetByUsername(request.Username)
	if err != nil {
		return errors.New("user not present")
	}
	if user.Token != mapper.EncryptAccessCode(request.OldPassword) {
		return errors.New("invalid token")
	}
	user.Token = mapper.EncryptAccessCode(request.NewPassword)
	return m.userStore.UpdateUser(user)
}
