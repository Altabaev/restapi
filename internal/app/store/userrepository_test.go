package store_test

import (
	"github.com/Altabaev/Go-Rest-Api/internal/app/model"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	_, err = s.User().Create(u)
	if err != nil {
		return
	}

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
