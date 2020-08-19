package model

import "testing"

func TestCreate(t *testing.T) {
	email := "test@email.com"
	password := "password"

	factory := &UserFactoryImplement{}

	user := factory.Create(email, password)
	if user.GetEmail() != "test@email.com" {
		t.Fatal("email is not matched")
	}

	if !user.ComparePassword("password") {
		t.Fatal("fail to compare password")
	}
}
