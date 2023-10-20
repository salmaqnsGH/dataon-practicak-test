package suubdivision

type Service interface {
	GetSubDivisions() ([]SubDivision, error)
	CreateSubDivision(input CreateSubDivisionInput, DivisionID int) (SubDivision, error)
	DeleteSubDivision(subDivisionID int) error
	UpdateSubDivision(inputID SubDivisionIDInput, inputData CreateSubDivisionInput) (SubDivision, error)
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

func (s *service) DeleteSubDivision(subDivisionID int) error {
	err := s.repository.Delete(subDivisionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateSubDivision(inputID SubDivisionIDInput, inputData CreateSubDivisionInput) (SubDivision, error) {
	subDivision, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return subDivision, err
	}

	subDivision.Name = inputData.Name

	updatedSubDivision, err := s.repository.Update(subDivision)
	if err != nil {
		return subDivision, err
	}

	return updatedSubDivision, nil
}
