package mapper

import (
	"testing"
	"time"

	"github.com/kyhsa93/rest-api-example/entity"
	"github.com/kyhsa93/rest-api-example/model"
)

func Test_EntityFromModel(t *testing.T) {
	now := time.Now()

	id := "userID"
	email := "test@email.com"
	password := "password"
	createdAt := now
	updatedAt := now

	user := model.UserImplement{}

	user.SetID(id)
	user.SetEmail(email)
	user.SetPassword(password)
	user.SetCreatedAt(createdAt)
	user.SetUpdatedAt(updatedAt)

	mapper := &UserMapperImplement{}

	if result := mapper.EntityFromModel(&user); result.ID != user.GetID() {
		t.Fail()
	}
	if result := mapper.EntityFromModel(&user); result.Email != user.GetEmail() {
		t.Fail()
	}
	if result := mapper.EntityFromModel(&user); result.Password != user.GetPassword() {
		t.Fail()
	}
	if result := mapper.EntityFromModel(&user); result.CreatedAt != user.GetCreatedAt() {
		t.Fail()
	}
	if result := mapper.EntityFromModel(&user); result.UpdatedAt != user.GetUpdatedAt() {
		t.Fail()
	}
}

func Test_ModelFromEntity(t *testing.T) {
	now := time.Now()
	entity := entity.UserEntity{
		ID:        "userID",
		Email:     "test@email.com",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mapper := &UserMapperImplement{}

	if result := mapper.ModelFromEntity(entity); result.GetID() != entity.ID {
		t.Fail()
	}
	if result := mapper.ModelFromEntity(entity); result.GetEmail() != entity.Email {
		t.Fail()
	}
	if result := mapper.ModelFromEntity(entity); result.GetPassword() == entity.Password {
		t.Fail()
	}
	if result := mapper.ModelFromEntity(entity); result.GetCreatedAt() != entity.CreatedAt {
		t.Fail()
	}
	if result := mapper.ModelFromEntity(entity); result.GetUpdatedAt() != entity.UpdatedAt {
		t.Fail()
	}
}
