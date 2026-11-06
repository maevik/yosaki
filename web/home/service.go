package home

type service struct {
	repository *repository
}

func newService(repo *repository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) PageTitle() string {
	return "Home"
}
