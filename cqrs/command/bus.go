package command

import (
	"errors"

	"github.com/kyhsa93/rest-api-example/command/create"
	"github.com/kyhsa93/rest-api-example/command/delete"
	"github.com/kyhsa93/rest-api-example/command/update"
)

// Bus handle command
type Bus interface {
	Handle(command interface{}) error
}

// BusImplement command bus implement
type BusImplement struct {
	createUserCommandHandler create.CommandHandler
	updateUserCommandHandler update.CommandHandler
	deleteUserCommandHandler delete.CommandHandler
}

// Handle call and execute command handlers by command type
func (bus *BusImplement) Handle(userCommand interface{}) error {
	switch userCommand := userCommand.(type) {
	case *create.Command:
		return bus.createUserCommandHandler.Handle(userCommand)
	case *update.Command:
		return bus.updateUserCommandHandler.Handle(userCommand)
	case *delete.Command:
		return bus.deleteUserCommandHandler.Handle(userCommand)
	default:
		return errors.New("can not handle command")
	}
}
