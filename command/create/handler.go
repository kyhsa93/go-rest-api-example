package create

import (
	"errors"

	"github.com/kyhsa93/rest-api-example/model"
	"github.com/kyhsa93/rest-api-example/repository"
)

// CommandHandler handler for Command
type CommandHandler interface {
	Handle(command *Command) error
}

// CommandHandlerImplement handler implement for Command
type CommandHandlerImplement struct {
	factory    model.UserFactory
	repository repository.UserRepository
}

// Handle handle Command
func (handler *CommandHandlerImplement) Handle(command *Command) error {
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

func (handler *CommandHandlerImplement) checkDuplicatedByEmail(email string) error {
	if user := handler.repository.FindByEmail(email); user.GetID() != "" {
		return errors.New("user email is already exists")
	}
	return nil
}
