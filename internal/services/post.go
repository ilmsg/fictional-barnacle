package services

import "github.com/ilmsg/fictional-barnacle/internal/models"

func GetPosts() []models.Post {
	user := GetUserById(1)
	posts := []models.Post{
		{Id: 1, Title: "Post 1", User: user},
		{Id: 2, Title: "Post 2", User: user},
		{Id: 3, Title: "Post 3", User: user},
	}
	return posts
}

func GetPostById(id int) models.Post {
	user := GetUserById(id)
	post := models.Post{
		Id: 3, Title: "Post 3", Content: "Post Content 3", User: user,
	}
	return post
}
