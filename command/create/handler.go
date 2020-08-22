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
	err := handler.checkDuplicatedByEmail(command.Email)
	if err != nil {
		return err
	}

	user := handler.factory.Create(command.Email, command.Password)

	handler.repository.Save(user)

	return nil
}

func (handler *CommandHandlerImplement) checkDuplicatedByEmail(emailString *string) error {
	email := model.Email(*emailString)
	if user := handler.repository.FindByEmail(&email); user.ID() == nil || *user.ID() != "" {
		return errors.New("user email is already exists")
	}
	return nil
}
