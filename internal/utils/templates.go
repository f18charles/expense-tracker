package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("web/template/pages/base.html", "web/template/pages/sidebar.html", "web/template/pages/"+page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func AuthTemplate(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("web/template/pages/authentication/base.html","web/template/pages/authentication/"+page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}