package model

import (
	"testing"
	"time"
)

func TestGetID(t *testing.T) {
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if id := user.GetID(); id != "userID" {
		t.Fail()
	}
}

func TestGetEmail(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updatedAt: now,
	}

	if email := user.GetEmail(); email != "test@email.com" {
		t.Fail()
	}
}

func TestGetPassword(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updatedAt: now,
	}

	if password := user.GetPassword(); password != "password" {
		t.Fail()
	}
}

func TestGetCreatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updatedAt: now,
	}
	if createdAt := user.GetCreatedAt(); createdAt != now {
		t.Fail()
	}
}

func TestGetUpdatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updatedAt: now,
	}
	if updatedAt := user.GetUpdatedAt(); updatedAt != now {
		t.Fail()
	}
}

func TestSetID(t *testing.T) {
	user := UserImplement{}
	user.SetID("userID")
	if userID := user.GetID(); userID != "userID" {
		t.Fail()
	}
}

func TestSetEmail(t *testing.T) {
	user := UserImplement{}
	user.SetEmail("test@email.com")
	if email := user.GetEmail(); email != "test@email.com" {
		t.Fail()
	}
}

func TestSetPassword(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updatedAt: now,
	}

	if err := user.SetPassword("new password"); err != nil {
		t.Fail()
	}
}

func TestSetCreatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{}
	user.SetCreatedAt(now)
	if createdAt := user.GetCreatedAt(); createdAt != now {
		t.Fail()
	}
}

func TestSetUpdatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{}
	user.SetUpdatedAt(now)
	if updatedAt := user.GetUpdatedAt(); updatedAt != now {
		t.Fail()
	}
}

func TestComparePassword(t *testing.T) {
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	user.SetPassword("new password")

	if compared := user.ComparePassword("new password"); !compared {
		t.Fail()
	}
}
