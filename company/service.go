package company

type Service interface {
	GetCompanies() ([]Company, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCompanies() ([]Company, error) {
	companies, err := s.repository.FindAll()
	if err != nil {
		return companies, err
	}

	return companies, nil
}
