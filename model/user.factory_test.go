package model

import "testing"

func TestCreate(t *testing.T) {
	emailString := "test@email.com"
	passwordString := "password"
	factory := &UserFactoryImplement{}

	user := factory.Create(&emailString, &passwordString)
	if user.Email() == nil || *user.Email() != Email("test@email.com") {
		t.Fatal("email is not matched")
	}

	password := Password("password")
	if user.Password() == nil || !*user.ComparePassword(&password) {
		t.Fatal("fail to compare password")
	}
}
