package model_test

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser(t *testing.T) {
	name := "Valdir"
	email := "valdir@test.com"
	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
}

func TestNewUser_WithoutName(t *testing.T) {
	email := "valdir@test.com"
	_, err := model.NewUser("", email)

	require.NotNil(t, err)
	require.Equal(t, "name: Missing required field", err.Error())
}

func TestNewUser_WithoutEmail(t *testing.T) {
	name := "Valdir"
	_, err := model.NewUser(name, "")

	require.NotNil(t, err)
	require.Equal(t, "email: Missing required field", err.Error())
}