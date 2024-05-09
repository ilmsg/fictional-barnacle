package handler

import (
	"net/http"

	"github.com/ilmsg/fictional-barnacle/internal/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{Title: "Home"}
	utils.RenderHTML(w, data, "layout", "navbar", "index")
}
