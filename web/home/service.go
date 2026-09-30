package home

type Service struct {
	repository *Repository
}

// TODO

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) PageTitle() string {
	return "Home"
}
