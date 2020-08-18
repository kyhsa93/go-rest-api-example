package model

import (
	"time"

	"github.com/google/uuid"
)

// UserFactory factory for user model
type UserFactory interface {
	Create(email, password string) User
}

// UserFactoryImplement factory implement for user model
type UserFactoryImplement struct{}

// Create create user
func (factory *UserFactoryImplement) Create(email, password string) User {
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	user := &UserImplement{}
	user.SetID(uuidV1.String())
	user.SetEmail(email)
	user.SetPassword(password)
	user.SetCreatedAt(time.Now())
	user.SetUpdatedAt(time.Now())

	return user
}
