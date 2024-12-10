package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-database-user/model"
	"testing"
)

type MockRoleRepository struct {
	mock.Mock
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) FindMstRole(ctx context.Context, roleId string) (model.MstRole, error) {
	called := m.Mock.Called(ctx, roleId)
	return called.Get(0).(model.MstRole), called.Error(1)
}

func (m *MockUserRepository) InsertUser(ctx context.Context, user model.MstUser) (model.MstUser, error) {
	called := m.Called(ctx, user)
	return called.Get(0).(model.MstUser), called.Error(1)
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	mockRoleRepo := new(MockRoleRepository)
	mockUserRepo := new(MockUserRepository)

	userServiceImpl := NewUserServiceImpl(mockUserRepo, mockRoleRepo)

	ctx := context.Background()

	inputUser := model.MstUser{
		Name:        "Supriadi Obo",
		Email:       "obo@gmail.com",
		Password:    "obotest",
		PhoneNumber: "731098419837",
	}

	expectedRole := model.MstRole{
		IdRole:   "ROLE002",
		RoleName: "User",
	}

	mockRoleRepo.On("FindMstRole", ctx, "ROLE002").Return(expectedRole, nil)

	expectedUser := model.MstUser{
		IdUser:      uuid.New().String(),
		Name:        "Supriadi Obo",
		Email:       "obo@gmail.com",
		Password:    "obotest",
		PhoneNumber: "731098419837",
		Role:        expectedRole,
	}

	mockUserRepo.On("InsertUser", ctx, mock.AnythingOfType("model.MstUser")).Return(expectedUser, nil)

	createUser := userServiceImpl.CreateUser(ctx, inputUser)

	assert.NotNil(t, createUser)
	assert.Equal(t, expectedUser.Name, createUser.Name)
	assert.Equal(t, expectedRole, createUser.Role)

	mockRoleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
