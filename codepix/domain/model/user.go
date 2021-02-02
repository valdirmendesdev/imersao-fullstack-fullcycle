package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID    string `json:"id" valid:"uuid"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		ID:    uuid.NewV4().String(),
		Name:  name,
		Email: email,
	}

	if err := user.isValid(); err != nil {
		return nil, err
	}

	return &user, nil
}
