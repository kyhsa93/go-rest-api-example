package update

import (
	"testing"

	"github.com/kyhsa93/rest-api-example/model"
)

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

func (repository *RepositoryMock) Delete(user model.User) {}

func TestHandle(t *testing.T) {
	command := Command{UserID: "userID", Password: "new password"}

	handler := CommandHandlerImplement{
		repository: &RepositoryMock{},
	}

	if err := handler.Handle(&command); err != nil {
		t.Fatal(err)
	}
}
