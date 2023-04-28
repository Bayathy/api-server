package db

import (
	"github.com/bayathy/api-server/cmd/entity"
	"github.com/bayathy/api-server/graph/model"
)

func ConvertUser(e *entity.User) *model.User {
	return &model.User{
		UUID: e.Uuid,
		ID:   e.ID,
	}
}

func ConvertArticle(e *entity.Article) *model.Article {
	return &model.Article{
		ID:        e.Id,
		Title:     e.Title,
		URL:       e.Url,
		UserID:    e.UserId,
		CreatedAt: e.CreatedAt,
		Done:      e.Done,
	}
}
