package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	if err := templates.ExecuteTemplate(w, "layout", data); err != nil {
		panic(err)
	}
}
