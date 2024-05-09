package handler

import (
	"net/http"

	"github.com/ilmsg/fictional-barnacle/internal/models"
	"github.com/ilmsg/fictional-barnacle/internal/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	data := &models.Profile{
		Name: "Scott Tiger",
		Bio:  "FullStack Developer",
	}
	utils.RenderHTML(w, data, "layout", "navbar", "users/profile")
}

func Register(w http.ResponseWriter, r *http.Request) {
	data := ""
	utils.RenderHTML(w, data, "layout", "navbar", "users/register")
}

func Login(w http.ResponseWriter, r *http.Request) {
	data := ""
	utils.RenderHTML(w, data, "layout", "navbar", "users/login")
}

func Setting(w http.ResponseWriter, r *http.Request) {
	data := ""
	utils.RenderHTML(w, data, "layout", "navbar", "users/setting")
}
func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
