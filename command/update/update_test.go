package update

import (
	"testing"
	"time"

	"github.com/kyhsa93/rest-api-example/model"
)

type RepositoryMock struct{}

func (repository *RepositoryMock) Save(user model.User) {}

func (repository *RepositoryMock) FindByID(userID *model.UserID) *model.UserImplement {
	email := model.Email("test@email.com")
	password := model.Password("password")
	createdAt := model.CreatedAt(time.Now())
	updatedAt := model.UpdatedAt(time.Now())

	user := &model.UserImplement{}
	user.SetID(userID)
	user.SetEmail(&email)
	user.SetPassword(&password)
	user.SetCreatedAt(&createdAt)
	user.SetUpdatedAt(&updatedAt)
	return user
}

func (repository *RepositoryMock) FindByEmail(email *model.Email) *model.UserImplement {
	userID := model.UserID("userID")
	password := model.Password("password")
	createdAt := model.CreatedAt(time.Now())
	updatedAt := model.UpdatedAt(time.Now())

	user := &model.UserImplement{}
	user.SetID(&userID)
	user.SetPassword(&password)
	user.SetCreatedAt(&createdAt)
	user.SetUpdatedAt(&updatedAt)
	return user
}

func (repository *RepositoryMock) Delete(user model.User) {}

func TestHandle(t *testing.T) {
	userID := "userID"
	password := "password"
	command := Command{UserID: &userID, Password: &password}

	handler := CommandHandlerImplement{
		repository: &RepositoryMock{},
	}

	if err := handler.Handle(&command); err != nil {
		t.Fatal(err)
	}
}
