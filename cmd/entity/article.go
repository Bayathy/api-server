package entity

type Article struct {
	ID    uint
	Title string
	Url   string
	Done  bool
	user  User
}
