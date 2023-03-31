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

func ConvertArticle(e *entity.Article) *model.Article {
	return &model.Article{
		Title: e.Title,
		URL:   e.Url,
		Done:  e.Done,
	}
}
