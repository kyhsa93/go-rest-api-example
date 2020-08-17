package model

import (
	"testing"
	"time"
)

func Test_GetID(t *testing.T) {
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: time.Now(),
		updateAt:  time.Now(),
	}

	if id := user.GetID(); id != "userID" {
		t.Fail()
	}
}

func Test_GetEmail(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updateAt:  now,
	}

	if email := user.GetEmail(); email != "test@email.com" {
		t.Fail()
	}
}

func Test_GetPassword(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updateAt:  now,
	}

	if password := user.GetPassword(); password != "password" {
		t.Fail()
	}
}

func Test_GetCreatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updateAt:  now,
	}
	if createdAt := user.GetCreatedAt(); createdAt != now {
		t.Fail()
	}
}

func Test_GetUpdatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updateAt:  now,
	}
	if updatedAt := user.GetUpdatedAt(); updatedAt != now {
		t.Fail()
	}
}

func Test_SetID(t *testing.T) {
	user := UserImplement{}
	user.SetID("userID")
	if userID := user.GetID(); userID != "userID" {
		t.Fail()
	}
}

func Test_SetEmail(t *testing.T) {
	user := UserImplement{}
	user.SetEmail("test@email.com")
	if email := user.GetEmail(); email != "test@email.com" {
		t.Fail()
	}
}

func Test_SetPassword(t *testing.T) {
	now := time.Now()
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: now,
		updateAt:  now,
	}

	if err := user.SetPassword("new password"); err != nil {
		t.Fail()
	}
}

func Test_SetCreatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{}
	user.SetCreatedAt(now)
	if createdAt := user.GetCreatedAt(); createdAt != now {
		t.Fail()
	}
}

func Test_SetUpdatedAt(t *testing.T) {
	now := time.Now()
	user := UserImplement{}
	user.SetUpdatedAt(now)
	if updatedAt := user.GetUpdatedAt(); updatedAt != now {
		t.Fail()
	}
}

func Test_ComparePassword(t *testing.T) {
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: time.Now(),
		updateAt:  time.Now(),
	}

	user.SetPassword("new password")

	if compared := user.ComparePassword("new password"); !compared {
		t.Fail()
	}
}
