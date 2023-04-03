// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	User      *User     `json:"user"`
}

type ArticleInput struct {
	UUID string `json:"uuid"`
}

type NewArticle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	UUID  string `json:"uuid"`
}

type NewUser struct {
	UUID string `json:"uuid"`
}

type User struct {
	ID      int        `json:"id"`
	UUID    string     `json:"uuid"`
	Article []*Article `json:"article,omitempty"`
}

type UserInput struct {
	UUID string `json:"uuid"`
}
