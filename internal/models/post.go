package models

type Post struct {
	Id      int
	Title   string
	Content string
	User    *User
}
