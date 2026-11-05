package view

import (
	"html/template"
	"net/http"
)

type Data struct {
	Page    string
	Code    int
	Message string
}

type Renderer struct {
	templates *template.Template
}

func NewRenderer(templates *template.Template) *Renderer {
	return &Renderer{templates: templates}
}

func (rdr *Renderer) Render(w http.ResponseWriter, name string, data Data) {
	rdr.RenderStatus(w, http.StatusOK, name, data)
}

func (rdr *Renderer) RenderStatus(w http.ResponseWriter, status int, name string, data Data) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := rdr.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}

func (rdr *Renderer) RenderError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	if err := rdr.templates.ExecuteTemplate(w, "error", Data{
		Page:    "Not Found",
		Code:    http.StatusNotFound,
		Message: "The page you are trying to reach does not exist.",
	}); err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}
