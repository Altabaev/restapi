package teststore_test

import (
	"github.com/Altabaev/Go-Rest-Api/internal/app/model"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	err = s.User().Create(u)
	if err != nil {
		return
	}

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
