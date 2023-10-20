package company

import (
	"dataon/division"
	"dataon/executiveCommittee"
	subdivision "dataon/subdivision"
)

type Service interface {
	GetCompanies() ([]Company, error)
	CreateCompany(input CreateCompanyInput) (Company, error)
	GetCompanyByID(inputID int) (CompanyOutput, error)
}

type service struct {
	repository            Repository
	ECRepository          executiveCommittee.Repository
	divisionRepository    division.Repository
	subDivisionRepository subdivision.Repository
}

func NewService(repository Repository, ECRepository executiveCommittee.Repository, divisionRepository division.Repository, subDivisionRepository subdivision.Repository) *service {
	return &service{repository, ECRepository, divisionRepository, subDivisionRepository}
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

func (s *service) GetCompanyByID(inputID int) (CompanyOutput, error) {
	var companyOutput CompanyOutput
	company, err := s.repository.FindByID(inputID)
	if err != nil {
		return companyOutput, err
	}

	companyOutput = CompanyOutput{
		ID:                 company.ID,
		Name:               company.Name,
		Executivecommittee: []ECOutput{}, // Initialize to an empty slice
	}

	executiveCommittes, err := s.ECRepository.FindAllByCompanyID(company.ID)
	if err != nil {
		return companyOutput, err
	}

	for _, ec := range executiveCommittes {
		ecOutput := ECOutput{
			ID:   ec.ID,
			Name: ec.Name,
		}

		divisions, err := s.divisionRepository.FindByECID(ec.ID)
		if err != nil {
			return companyOutput, err
		}

		for _, div := range divisions {
			divisionOutput := Division{
				ID:   div.ID,
				Name: div.Name,
			}

			subDivisions, err := s.subDivisionRepository.FindAllByDivisionID(div.ID)
			if err != nil {
				return companyOutput, err
			}

			for _, subDiv := range subDivisions {
				divisionOutput.SubDivision = append(divisionOutput.SubDivision, subDiv)
			}

			ecOutput.Division = append(ecOutput.Division, divisionOutput)
		}

		companyOutput.Executivecommittee = append(companyOutput.Executivecommittee, ecOutput)
	}

	return companyOutput, nil
}
