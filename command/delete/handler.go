package delete

import "github.com/kyhsa93/rest-api-example/repository"

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
	userID := command.UserID
	user := handler.repository.FindByID(userID)
	handler.repository.Delete(user)
	return nil
}
