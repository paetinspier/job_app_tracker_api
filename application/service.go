package application

type Service struct {
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repository: repo}
}

func (s *Service) GetApplicationById(id int) (*Application, error) {
	return s.Repository.FindApplicationById(id)
}
