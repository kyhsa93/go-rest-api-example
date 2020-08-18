package command

import (
	"testing"
	"time"

	"github.com/kyhsa93/rest-api-example/model"
)

type FactoryMock struct{}

func (factory *FactoryMock) Create(email, password string) model.User {
	user := &model.UserImplement{}
	user.SetID("userID")
	user.SetEmail(email)
	user.SetPassword(password)
	user.SetCreatedAt(time.Now())
	user.SetUpdatedAt(time.Now())
	return user
}

type RepositoryMock struct{}

func (repository *RepositoryMock) Save(user model.User) {}

func (repository *RepositoryMock) FindByID(userID string) model.User {
	user := &model.UserImplement{}
	user.SetID(userID)
	return user
}

func (repository *RepositoryMock) FindByEmail(email string) model.User {
	user := &model.UserImplement{}
	user.SetEmail(email)
	return user
}

func Test_Handle(t *testing.T) {
	command := CreateUserCommand{
		Email:    "test.@email.com",
		Password: "password",
	}

	handler := CreateUserCommandHandlerImplement{
		factory:    &FactoryMock{},
		repository: &RepositoryMock{},
	}

	if err := handler.Handle(command); err != nil {
		t.Fatal("can not handle CreateUserCommand")
	}
}
