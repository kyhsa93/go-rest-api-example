package mapper

import (
	"github.com/kyhsa93/rest-api-example/entity"
	"github.com/kyhsa93/rest-api-example/model"
)

// UserMapper user mapper for convert to between model and entity
type UserMapper interface {
	EntityFromModel(model model.User) *entity.UserEntity
	ModelFromEntity(entity entity.UserEntity) model.User
}

// UserMapperImplement mapper implement for user mapper
type UserMapperImplement struct{}

// EntityFromModel create entity from model
func (mapper *UserMapperImplement) EntityFromModel(model model.User) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        model.GetID(),
		Email:     model.GetEmail(),
		Password:  model.GetPassword(),
		CreatedAt: model.GetCreatedAt(),
		UpdatedAt: model.GetUpdatedAt(),
	}
}

// ModelFromEntity create model from entity
func (mapper *UserMapperImplement) ModelFromEntity(entity entity.UserEntity) *model.UserImplement {
	user := &model.UserImplement{}
	user.SetID(entity.ID)
	user.SetEmail(entity.Email)
	user.SetPassword(entity.Password)
	user.SetCreatedAt(entity.CreatedAt)
	user.SetUpdatedAt(entity.UpdatedAt)
	return user
}
