package suubdivision

type Service interface {
	GetSubDivisions() ([]SubDivision, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetSubDivisions() ([]SubDivision, error) {
	subDivisions, err := s.repository.FindAll()
	if err != nil {
		return subDivisions, err
	}

	return subDivisions, nil
}
