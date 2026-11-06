package auth

import "net/http"

type service struct {
	repository *repository
}

func newService(repo *repository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) IsAuthenticated(r *http.Request) bool {
	return s.repository.IsAuthenticated(r)
}

func (s *service) Signin(w http.ResponseWriter) {
	s.repository.SetAuthenticated(w)
}

func (s *service) Signup(w http.ResponseWriter) {
	s.repository.SetAuthenticated(w)
}
