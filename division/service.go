package division

type Service interface {
	GetDivisions() ([]Division, error)
	CreateDivision(input CreateDivisionInput, executivecommitteeID int) (Division, error)
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

func (s *service) CreateDivision(input CreateDivisionInput, executivecommitteeID int) (Division, error) {
	division := Division{}
	division.Name = input.Name
	division.ExecutiveCommitteeID = executivecommitteeID

	newDivision, err := s.repository.Save(division)
	if err != nil {
		return newDivision, err
	}

	return newDivision, nil
}
