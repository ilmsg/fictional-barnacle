package services

import "github.com/ilmsg/fictional-barnacle/internal/models"

func GetUserById(id int) *models.User {
	user := &models.User{Email: "scott@gmail.com"}
	return user
}
