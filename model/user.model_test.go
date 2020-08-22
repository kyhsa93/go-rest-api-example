package model

import (
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TestID(t *testing.T) {
	userID := UserID("userID")
	user := UserImplement{id: &userID}

	if id := user.ID(); *id != userID {
		t.Fatal()
	}
}

func TestEmail(t *testing.T) {
	email := Email("test@email.com")
	user := UserImplement{email: &email}

	if user.Email() == nil || *user.Email() != email {
		t.Fatal()
	}
}

func TestPassword(t *testing.T) {
	password := Password("password")
	user := UserImplement{password: &password}

	if *user.Password() != password {
		t.Fatal()
	}
}

func TestCreatedAt(t *testing.T) {
	createdAt := CreatedAt(time.Now())
	user := UserImplement{createdAt: &createdAt}

	if *user.CreatedAt() != createdAt {
		t.Fatal()
	}
}

func TestUpdatedAt(t *testing.T) {
	updatedAt := UpdatedAt(time.Now())
	user := UserImplement{updatedAt: &updatedAt}

	if *user.UpdatedAt() != updatedAt {
		t.Fatal()
	}
}

func TestSetID(t *testing.T) {
	userID := UserID("userID")
	user := UserImplement{}
	user.SetID(&userID)

	if *user.id != userID {
		t.Fatal()
	}
}

func TestSetEmail(t *testing.T) {
	email := Email("test@email.com")
	user := UserImplement{}
	user.SetEmail(&email)

	if *user.email != email {
		t.Fatal()
	}
}

func TestSetPassword(t *testing.T) {
	password := Password("password")

	user := UserImplement{}
	user.SetPassword(&password)

	if user.password == nil || *user.password == Password("password") {
		t.Fatal()
	}
}

func TestSetCreatedAt(t *testing.T) {
	user := UserImplement{}
	createdAt := CreatedAt(time.Now())
	user.SetCreatedAt(&createdAt)

	if *user.createdAt != createdAt {
		t.Fatal()
	}
}

func TestSetUpdatedAt(t *testing.T) {
	user := UserImplement{}
	updatedAt := UpdatedAt(time.Now())
	user.SetUpdatedAt(&updatedAt)

	if *user.updatedAt != updatedAt {
		t.Fatal()
	}
}

func TestComparePassword(t *testing.T) {
	newPassword := Password("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}

	password := Password(string(hash))
	user := UserImplement{password: &password}
	if !*user.ComparePassword(&newPassword) {
		t.Fatal()
	}
}
