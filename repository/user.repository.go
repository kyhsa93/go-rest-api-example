package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/rest-api-example/entity"
	"github.com/kyhsa93/rest-api-example/model"
)

// UserRepository user repository
type UserRepository interface {
	Save(user model.User)
	FindByID(userID string) model.User
	FindByEmail(email string) model.User
	Delete(user model.User)
}

// UserRepositoryImplement implement for user repository
type UserRepositoryImplement struct {
	database *gorm.DB
}

// Save insert or update given user model data
func (repository *UserRepositoryImplement) Save(user model.User) {
	entity := entityFromModel(user)
	repository.database.Save(entity)
}

// FindByID find single user data by ID
func (repository *UserRepositoryImplement) FindByID(userID string) model.User {
	user := entity.UserEntity{}
	repository.database.First(&user, userID)
	return modelFromEntity(user)
}

// FindByEmail find single user data by email
func (repository *UserRepositoryImplement) FindByEmail(email string) model.User {
	user := entity.UserEntity{}
	repository.database.Where(&entity.UserEntity{Email: email}).First(&user)
	return modelFromEntity(user)
}

// Delete delete user data
func (repository *UserRepositoryImplement) Delete(user model.User) {
	entity := entityFromModel(user)
	repository.database.Delete(entity)
}

func entityFromModel(model model.User) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        model.GetID(),
		Email:     model.GetEmail(),
		Password:  model.GetPassword(),
		CreatedAt: model.GetCreatedAt(),
		UpdatedAt: model.GetUpdatedAt(),
	}
}

func modelFromEntity(entity entity.UserEntity) model.User {
	user := &model.UserImplement{}
	user.SetID(entity.ID)
	user.SetEmail(entity.Email)
	user.SetPassword(entity.Password)
	user.SetCreatedAt(entity.CreatedAt)
	user.SetUpdatedAt(entity.UpdatedAt)
	return user
}
