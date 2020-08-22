package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserID uuid for user identifier
type UserID string

// Email email
type Email string

// Password user password hash
type Password string

// CreatedAt datetime when user created
type CreatedAt time.Time

// UpdatedAt datetime when user last updated
type UpdatedAt time.Time

// User domain model for user
type User interface {
	ID() *UserID
	Email() *Email
	Password() *Password
	CreatedAt() *CreatedAt
	UpdatedAt() *UpdatedAt
	SetID(id *UserID)
	SetEmail(email *Email)
	SetPassword(password *Password) error
	SetCreatedAt(datetime *CreatedAt)
	SetUpdatedAt(datetime *UpdatedAt)
	ComparePassword(password *Password) *bool
}

// UserImplement domain model implement for user
type UserImplement struct {
	id        *UserID
	email     *Email
	password  *Password
	createdAt *CreatedAt
	updatedAt *UpdatedAt
}

// ID get id
func (user *UserImplement) ID() *UserID {
	return user.id
}

// Email get email
func (user *UserImplement) Email() *Email {
	return user.email
}

// Password get password
func (user *UserImplement) Password() *Password {
	return user.password
}

// CreatedAt get created datetime
func (user *UserImplement) CreatedAt() *CreatedAt {
	return user.createdAt
}

// UpdatedAt get updated datetime
func (user *UserImplement) UpdatedAt() *UpdatedAt {
	return user.updatedAt
}

// SetID set id
func (user *UserImplement) SetID(id *UserID) {
	user.id = id
}

// SetEmail set email
func (user *UserImplement) SetEmail(email *Email) {
	user.email = email
}

// SetPassword set user password
func (user *UserImplement) SetPassword(password *Password) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	newPassword := Password(string(hash))
	user.password = &newPassword
	return nil
}

// SetCreatedAt set createdAt
func (user *UserImplement) SetCreatedAt(datetime *CreatedAt) {
	user.createdAt = datetime
}

// SetUpdatedAt set updatedAt
func (user *UserImplement) SetUpdatedAt(datetime *UpdatedAt) {
	user.updatedAt = datetime
}

// ComparePassword compare password
func (user *UserImplement) ComparePassword(password *Password) *bool {
	err := bcrypt.CompareHashAndPassword([]byte(*user.password), []byte(*password))
	result := err == nil
	log.Println(result)
	return &result
}
