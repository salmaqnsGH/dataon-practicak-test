package executiveCommittee

type Service interface {
	GetExecutivecommitties() ([]Executivecommittee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetExecutivecommitties() ([]Executivecommittee, error) {
	Executivecommitties, err := s.repository.FindAll()
	if err != nil {
		return Executivecommitties, err
	}

	return Executivecommitties, nil
}
