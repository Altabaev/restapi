package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
