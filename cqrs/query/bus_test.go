package query

import (
	"errors"
	"testing"

	"github.com/kyhsa93/rest-api-example/model"
	"github.com/kyhsa93/rest-api-example/query"
)

type findUserQueryHandlerMock struct{}

func (handler *findUserQueryHandlerMock) Handle(query *query.FindUserQuery) (model.User, error) {
	return nil, errors.New("findUserQueryHandler")
}

func TestHandle(t *testing.T) {
	bus := BusImplement{
		findUserQueryHandler: &findUserQueryHandlerMock{},
	}

	findUserQuery := query.FindUserQuery{}
	if _, err := bus.Handle(&findUserQuery); err.Error() != "findUserQueryHandler" {
		t.Fatal("queryBus can not call find user query handler")
	}
}
