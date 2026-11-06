package auth

import "net/http"

type repository struct{}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) IsAuthenticated(req *http.Request) bool {
	cookie, err := req.Cookie("authenticated")
	return err == nil && cookie.Value == "true"
}

func (r *repository) SetAuthenticated(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "authenticated",
		Value:    "true",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
