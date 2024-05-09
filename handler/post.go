package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilmsg/fictional-barnacle/internal/models"
	"github.com/ilmsg/fictional-barnacle/internal/services"
	"github.com/ilmsg/fictional-barnacle/internal/utils"
)

func PostList(w http.ResponseWriter, r *http.Request) {
	posts := services.GetPosts()
	data := struct{ Posts []models.Post }{Posts: posts}

	utils.RenderHTML(w, data, "layout", "navbar", "posts/list")
}

func PostDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramId := vars["id"]

	id, _ := strconv.Atoi(paramId)
	fmt.Println("post-detail id:", id)

	post := services.GetPostById(id)
	data := struct{ Post models.Post }{Post: post}

	utils.RenderHTML(w, data, "layout", "navbar", "posts/detail")
}
