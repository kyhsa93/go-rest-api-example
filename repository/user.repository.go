package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/rest-api-example/mapper"
	"github.com/kyhsa93/rest-api-example/model"
)

// UserRepository user repository
type UserRepository interface {
	Save(user model.User)
}

// UserRepositoryImplement implement for user repository
type UserRepositoryImplement struct {
	mapper   mapper.UserMapper
	database *gorm.DB
}

// Save insert or update given user model data
func (repository *UserRepositoryImplement) Save(user model.User) {
	entity := repository.mapper.EntityFromModel(user)
	repository.database.Save(entity)
}
