package model

import (
	"time"

	"github.com/google/uuid"
)

// UserFactory factory for user model
type UserFactory interface {
	Create(emailString, passwordString *string) *UserImplement
}

// UserFactoryImplement factory implement for user model
type UserFactoryImplement struct{}

// Create create user
func (factory *UserFactoryImplement) Create(emailString, passwordString *string) *UserImplement {
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	userID := UserID(uuidV1.String())
	email := Email(*emailString)
	password := Password(*passwordString)
	createdAt := CreatedAt(time.Now())
	updatedAt := UpdatedAt(time.Now())

	user := &UserImplement{}
	user.SetID(&userID)
	user.SetEmail(&email)
	if err = user.SetPassword(&password); err != nil {
		panic(err)
	}
	user.SetCreatedAt(&createdAt)
	user.SetUpdatedAt(&updatedAt)

	return user
}
