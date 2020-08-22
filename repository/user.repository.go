package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/rest-api-example/entity"
	"github.com/kyhsa93/rest-api-example/model"
)

// UserRepository user repository
type UserRepository interface {
	Save(user model.User)
	Delete(user model.User)
	FindByID(userID *model.UserID) *model.UserImplement
	FindByEmail(email *model.Email) *model.UserImplement
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

// Delete delete user data
func (repository *UserRepositoryImplement) Delete(user model.User) {
	entity := entityFromModel(user)
	repository.database.Delete(entity)
}

// FindByID find single user data by ID
func (repository *UserRepositoryImplement) FindByID(userID *model.UserID) *model.UserImplement {
	user := entity.UserEntity{}
	repository.database.First(&user, userID)
	return modelFromEntity(user)
}

// FindByEmail find single user data by email
func (repository *UserRepositoryImplement) FindByEmail(email *model.Email) *model.UserImplement {
	user := entity.UserEntity{}
	repository.database.Where(&entity.UserEntity{Email: string(*email)}).First(&user)
	return modelFromEntity(user)
}

func entityFromModel(model model.User) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        string(*model.ID()),
		Email:     string(*model.Email()),
		Password:  string(*model.Password()),
		CreatedAt: time.Time(*model.CreatedAt()),
		UpdatedAt: time.Time(*model.UpdatedAt()),
	}
}

func modelFromEntity(entity entity.UserEntity) *model.UserImplement {
	userID := model.UserID(entity.ID)
	email := model.Email(entity.Email)
	password := model.Password(entity.Password)
	createdAt := model.CreatedAt(entity.CreatedAt)
	updatedAt := model.UpdatedAt(entity.UpdatedAt)

	user := &model.UserImplement{}
	user.SetID(&userID)
	user.SetEmail(&email)
	user.SetPassword(&password)
	user.SetCreatedAt(&createdAt)
	user.SetUpdatedAt(&updatedAt)
	return user
}
