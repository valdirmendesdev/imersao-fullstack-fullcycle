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