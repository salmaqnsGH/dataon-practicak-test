package suubdivision

type Service interface {
	GetSubDivisions() ([]SubDivision, error)
	CreateSubDivision(input CreateSubDivisionInput, DivisionID int) (SubDivision, error)
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

func (s *service) CreateSubDivision(input CreateSubDivisionInput, DivisionID int) (SubDivision, error) {
	subDivision := SubDivision{}
	subDivision.Name = input.Name
	subDivision.DivisionID = DivisionID

	newSubDivision, err := s.repository.Save(subDivision)
	if err != nil {
		return newSubDivision, err
	}

	return newSubDivision, nil
}
