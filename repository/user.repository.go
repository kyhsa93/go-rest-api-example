package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/rest-api-example/entity"
	"github.com/kyhsa93/rest-api-example/mapper"
	"github.com/kyhsa93/rest-api-example/model"
)

// UserRepository user repository
type UserRepository interface {
	Save(user model.User)
	FindByID(userID string) model.User
	FindByEmail(email string) model.User
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

// FindByID find single user data by ID
func (repository *UserRepositoryImplement) FindByID(userID string) model.User {
	user := entity.UserEntity{}
	repository.database.First(&user, userID)
	return repository.mapper.ModelFromEntity(user)
}

// FindByEmail find single user data by email
func (repository *UserRepositoryImplement) FindByEmail(email string) model.User {
	user := entity.UserEntity{}
	repository.database.Where(&entity.UserEntity{Email: email}).First(&user)
	return repository.mapper.ModelFromEntity(user)
}
