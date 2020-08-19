package query

import (
	"testing"
	"time"

	"github.com/kyhsa93/rest-api-example/model"
)

type RepositoryMock struct{}

func (repository *RepositoryMock) Save(user model.User) {}

func (repository *RepositoryMock) FindByID(userID string) model.User {
	user := &model.UserImplement{}
	user.SetID("userID")
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
	query := FindUserQuery{UserID: "userID"}

	handler := FindUserQueryHandlerImplement{
		repository: &RepositoryMock{},
	}

	if _, err := handler.Handle(&query); err != nil {
		t.Fatal(err)
	}
}
