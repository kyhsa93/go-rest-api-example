package model

import (
	"testing"
	"time"
)

func Test_SetPassword(t *testing.T) {
	user := UserImplement{
		id:        "userID",
		email:     "test@email.com",
		password:  "password",
		createdAt: time.Now(),
		updateAt:  time.Now(),
	}

	if err := user.SetPassword("new password"); err != nil {
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
