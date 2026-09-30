package auth

import "net/http"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO

func (r *Repository) IsAuthenticated(req *http.Request) bool {
	cookie, err := req.Cookie("authenticated")
	return err == nil && cookie.Value == "true"
}

func (r *Repository) SetAuthenticated(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "authenticated",
		Value:    "true",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
