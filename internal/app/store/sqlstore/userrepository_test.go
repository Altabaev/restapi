package sqlstore_test

import (
	"github.com/Altabaev/Go-Rest-Api/internal/app/model"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store"
	"github.com/Altabaev/Go-Rest-Api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
	err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)
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

func TestUserRepository_Find(t *testing.T) {

	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	_ = s.User().Create(u)

	u, err := s.User().Find(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
