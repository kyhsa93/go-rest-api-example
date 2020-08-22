package update

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
	repository repository.UserRepository
}

// Handle handle Command
func (handler *CommandHandlerImplement) Handle(command *Command) error {
	userID := model.UserID(*command.UserID)
	password := model.Password(*command.Password)

	user := handler.repository.FindByID(&userID)
	if *user.ID() == "" {
		return errors.New("user not found")
	}

	if err := user.SetPassword(&password); err != nil {
		return err
	}

	handler.repository.Save(user)

	return nil
}
