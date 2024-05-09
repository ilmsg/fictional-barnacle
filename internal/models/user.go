package models

type User struct {
	Email    string
	Password string
	Profile  *Profile
}
