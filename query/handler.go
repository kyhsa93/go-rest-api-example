package query

import (
	"errors"

	"github.com/kyhsa93/rest-api-example/model"
	"github.com/kyhsa93/rest-api-example/repository"
)

// FindUserQueryHandler handler for FindUserQuery
type FindUserQueryHandler interface {
	Handle(query *FindUserQuery) (model.User, error)
}

// FindUserQueryHandlerImplement handler implement for FindUserQuery
type FindUserQueryHandlerImplement struct {
	repository repository.UserRepository
}

// Handle handle FindUserQuery
func (handler *FindUserQueryHandlerImplement) Handle(query *FindUserQuery) (model.User, error) {
	userID := model.UserID(query.UserID)
	user := handler.repository.FindByID(&userID)

	if *user.ID() == "" {
		return nil, errors.New("user not found")
	}

	return user, nil
}
