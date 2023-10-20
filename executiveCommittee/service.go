package executiveCommittee

type Service interface {
	GetExecutivecommitties() ([]Executivecommittee, error)
	CreateExecutivecommittee(input CreateExecutiveCommitteeInput, companyID int) (Executivecommittee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetExecutivecommitties() ([]Executivecommittee, error) {
	executivecommitties, err := s.repository.FindAll()
	if err != nil {
		return executivecommitties, err
	}

	return executivecommitties, nil
}

func (s *service) CreateExecutivecommittee(input CreateExecutiveCommitteeInput, companyID int) (Executivecommittee, error) {
	executivecommittee := Executivecommittee{}
	executivecommittee.Name = input.Name
	executivecommittee.CompanyID = companyID

	newexecutivecommittee, err := s.repository.Save(executivecommittee)
	if err != nil {
		return newexecutivecommittee, err
	}

	return newexecutivecommittee, nil
}
