package db

import (
	"github.com/bayathy/api-server/cmd/entity"
	"github.com/bayathy/api-server/graph/model"
)

func ConvertUser(e *entity.User) *model.User {
	return &model.User{
		UUID: e.Uuid,
	}
}
