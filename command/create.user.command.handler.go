package command

import (
	"errors"

	"github.com/kyhsa93/rest-api-example/model"
	"github.com/kyhsa93/rest-api-example/repository"
)

// CreateUserCommandHandler handler for CreateUserCommand
type CreateUserCommandHandler interface {
	Handle(command CreateUserCommand)
}

// CreateUserCommandHandlerImplement handler implement for CreateUSerCommand
type CreateUserCommandHandlerImplement struct {
	factory    model.UserFactory
	repository repository.UserRepository
}

// Handle handle CreateUserCommand
func (handler *CreateUserCommandHandlerImplement) Handle(command CreateUserCommand) error {
	email := command.Email
	password := command.Password

	err := handler.checkDuplicatedByEmail(email)
	if err != nil {
		return err
	}

	user := handler.factory.Create(email, password)

	handler.repository.Save(user)

	return nil
}

func (handler *CreateUserCommandHandlerImplement) checkDuplicatedByEmail(email string) error {
	if user := handler.repository.FindByEmail(email); user.GetID() != "" {
		return errors.New("user email is already exists")
	}
	return nil
}
