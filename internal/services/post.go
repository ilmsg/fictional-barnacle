package services

import (
	"errors"
	"time"

	"github.com/ilmsg/fictional-barnacle/internal/models"
)

var users = map[int]*models.User{
	1: {Id: 1, Email: "", Password: "", Profile: &models.Profile{Id: 1, Name: "", Bio: ""}},
}

var posts = map[int]models.Post{
	1: {Id: 1, Title: "Post 1", Content: "Post Content 1", CreatedAt: time.Now().UTC(), User: users[1]},
	2: {Id: 2, Title: "Post 2", Content: "Post Content 2", CreatedAt: time.Now().UTC(), User: users[1]},
	3: {Id: 3, Title: "Post 3", Content: "Post Content 3", CreatedAt: time.Now().UTC(), User: users[1]},
}

func GetPosts() []models.Post {
	var allPosts []models.Post
	for _, post := range posts {
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func GetPostById(id int) (post models.Post, err error) {
	if _, ok := posts[id]; !ok {
		return models.Post{Id: 0, Title: "", Content: "", User: nil}, errors.New("post not found")
	}
	return posts[id], nil
}
