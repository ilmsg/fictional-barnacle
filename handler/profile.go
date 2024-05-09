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
