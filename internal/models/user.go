package models

import "time"

type User struct {
	ID        string
	Username  string
	Name      string
	Email     string
	Bio       string
	ImageUrl  string
	createdAt time.Time
}

type Post struct {
	Id        string
	Content   string
	AuthorId  string
	Likes     int
	CreatedAt time.Time
}

type Comment struct {
	ID        string
	Content   string
	PostId    string
	UserId    string
	likes     int
	CreatedAt time.Time
}
