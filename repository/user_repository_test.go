package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"
)

func TestInsertUser_Success(t *testing.T) {

	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	newUserRepositoryImpl := NewUserRepositoryImpl(sql)
	newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

	ctx := context.Background()

	role, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE001")

	mstUser := model.MstUser{
		IdUser:      uuid.NewString(),
		Name:        "Test",
		Email:       "test@test.com",
		Password:    "pass",
		PhoneNumber: "09347343",
		Role:        role,
	}

	insertUser, err := newUserRepositoryImpl.InsertUser(ctx, mstUser)

	if err != nil {
		panic(err)
	}

	assert.NotNil(t, insertUser)
	assert.Equal(t, mstUser, insertUser)
	assert.Equal(t, mstUser.IdUser, insertUser.IdUser)
	assert.Equal(t, mstUser.Email, insertUser.Email)
	assert.Equal(t, mstUser.Role, insertUser.Role)
}
