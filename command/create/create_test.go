package create

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
	user.SetEmail("test@email.com")
	user.SetPassword("password")
	user.SetCreatedAt(time.Now())
	user.SetUpdatedAt(time.Now())
	return user
}

func (repository *RepositoryMock) FindByEmail(email string) model.User {
	user := &model.UserImplement{}
	return user
}

func (repository *RepositoryMock) Delete(user model.User) {}

func TestHandle(t *testing.T) {
	command := Command{
		Email:    "test.@email.com",
		Password: "password",
	}

	handler := CommandHandlerImplement{
		factory:    &FactoryMock{},
		repository: &RepositoryMock{},
	}

	if err := handler.Handle(&command); err != nil {
		t.Fatal(err)
	}
}
