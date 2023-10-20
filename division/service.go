package division

type Service interface {
	GetDivisions() ([]Division, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetDivisions() ([]Division, error) {
	divisions, err := s.repository.FindAll()
	if err != nil {
		return divisions, err
	}

	return divisions, nil
}
