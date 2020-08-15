package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User domain model for user
type User interface {
	SetPassword(password string)
	ComparePassword(password string)
}

// UserImplement domain model implement for user
type UserImplement struct {
	id        string
	email     string
	password  string
	createdAt time.Time
	updateAt  time.Time
}

// SetPassword update user password
func (user *UserImplement) SetPassword(password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	user.password = string(hash)
	return nil
}

// ComparePassword compare password
func (user *UserImplement) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	return err == nil
}
