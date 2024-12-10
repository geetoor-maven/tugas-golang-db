package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-database-user/config"
	"testing"
)

func TestFindMstRole_Success(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

	ctx := context.Background()

	role, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE002")
	if err != nil {
		panic(err)
	}

	assert.NotNil(t, role)
	assert.NotEqual(t, "ROLE001", role.IdRole)
	assert.Equal(t, "ROLE002", role.IdRole)
	assert.Equal(t, "User", role.RoleName)
}

func TestFindMstRole_Fail(t *testing.T) {
	sql, err := config.OpenConnectionPostgresSQL()
	if err != nil {
		panic(err)
	}

	newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

	ctx := context.Background()

	mstRole, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE004")

	assert.Emptyf(t, "", mstRole.IdRole, "")
}
