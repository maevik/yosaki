package auth

import "net/http"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) IsAuthenticated(r *http.Request) bool {
	return s.repository.IsAuthenticated(r)
}

func (s *Service) Signin(w http.ResponseWriter) {
	s.repository.SetAuthenticated(w)
}

func (s *Service) Signup(w http.ResponseWriter) {
	s.repository.SetAuthenticated(w)
}
