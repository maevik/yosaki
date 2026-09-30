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

func (r *Renderer) Render(w http.ResponseWriter, name string, data Data) {
	r.RenderStatus(w, http.StatusOK, name, data)
}

func (r *Renderer) RenderStatus(w http.ResponseWriter, status int, name string, data Data) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := r.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
	}
}
