package company

type Service interface {
	GetCompanies() ([]Company, error)
	CreateCompany(input CreateCompanyInput) (Company, error)
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

func (s *service) CreateCompany(input CreateCompanyInput) (Company, error) {
	company := Company{}
	company.Name = input.Name

	newCompany, err := s.repository.Save(company)
	if err != nil {
		return newCompany, err
	}

	return newCompany, nil
}
