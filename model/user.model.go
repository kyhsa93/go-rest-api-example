package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User domain model for user
type User interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(id string)
	SetEmail(email string)
	SetPassword(password string) error
	SetCreatedAt(datetime time.Time)
	SetUpdatedAt(datetime time.Time)
	ComparePassword(password string) bool
}

// UserImplement domain model implement for user
type UserImplement struct {
	id        string
	email     string
	password  string
	createdAt time.Time
	updateAt  time.Time
}

// GetID get id
func (user *UserImplement) GetID() string {
	return user.id
}

// GetEmail get email
func (user *UserImplement) GetEmail() string {
	return user.email
}

// GetPassword get password
func (user *UserImplement) GetPassword() string {
	return user.password
}

// GetCreatedAt get created datetime
func (user *UserImplement) GetCreatedAt() time.Time {
	return user.createdAt
}

// GetUpdatedAt get updated datetime
func (user *UserImplement) GetUpdatedAt() time.Time {
	return user.updateAt
}

// SetID set id
func (user *UserImplement) SetID(id string) {
	user.id = id
}

// SetEmail set email
func (user *UserImplement) SetEmail(email string) {
	user.email = email
}

// SetPassword set user password
func (user *UserImplement) SetPassword(password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	user.password = string(hash)
	return nil
}

// SetCreatedAt set createdAt
func (user *UserImplement) SetCreatedAt(datetime time.Time) {
	user.createdAt = datetime
}

// SetUpdatedAt set updatedAt
func (user *UserImplement) SetUpdatedAt(datetime time.Time) {
	user.updateAt = datetime
}

// ComparePassword compare password
func (user *UserImplement) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	return err == nil
}
