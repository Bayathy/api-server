// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Done  bool   `json:"done"`
	User  *User  `json:"user,omitempty"`
}

type CreateArticleInput struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	UUID  string `json:"uuid"`
}

type CreateUserInput struct {
	UUID string `json:"uuid"`
}

type User struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
}