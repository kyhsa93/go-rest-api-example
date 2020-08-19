package query

import (
	"errors"

	"github.com/kyhsa93/rest-api-example/query"
)

// Bus handle query
type Bus interface {
	Handle(query interface{}) (interface{}, error)
}

// BusImplement query bu implement
type BusImplement struct {
	findUserQueryHandler query.FindUserQueryHandler
}

// Handle call and excute query handlers by query type
func (bus *BusImplement) Handle(userQuery interface{}) (interface{}, error) {
	switch userQuery := userQuery.(type) {
	case *query.FindUserQuery:
		return bus.findUserQueryHandler.Handle(userQuery)
	default:
		return nil, errors.New("can not handle query")
	}
}
