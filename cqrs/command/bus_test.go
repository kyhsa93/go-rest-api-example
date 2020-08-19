package command

import (
	"errors"
	"testing"

	"github.com/kyhsa93/rest-api-example/command/create"
	"github.com/kyhsa93/rest-api-example/command/delete"
	"github.com/kyhsa93/rest-api-example/command/update"
)

type createCommandHandlerMock struct{}

func (handler *createCommandHandlerMock) Handle(command *create.Command) error {
	return errors.New("createCommandHandler")
}

type updateCommandHandlerMock struct{}

func (handler *updateCommandHandlerMock) Handle(command *update.Command) error {
	return errors.New("updateCommandHandler")
}

type deleteCommandHandlerMock struct{}

func (handler *deleteCommandHandlerMock) Handle(command *delete.Command) error {
	return errors.New("deleteCommandHandler")
}

func TestHandle(t *testing.T) {
	bus := BusImplement{
		createUserCommandHandler: &createCommandHandlerMock{},
		updateUserCommandHandler: &updateCommandHandlerMock{},
		deleteUserCommandHandler: &deleteCommandHandlerMock{},
	}

	createCommand := create.Command{}
	if bus.Handle(&createCommand).Error() != "createCommandHandler" {
		t.Fatal("commandBus can not call create command handler")
	}

	updateCommand := update.Command{}
	if bus.Handle(&updateCommand).Error() != "updateCommandHandler" {
		t.Fatal("commandBus can not call update command handler")
	}

	deleteCommand := delete.Command{}
	if bus.Handle(&deleteCommand).Error() != "deleteCommandHandler" {
		t.Fatal("commandBus can not call delete command handler")
	}
}
