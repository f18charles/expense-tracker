package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("web/template/base.html", "web/template/"+page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}
