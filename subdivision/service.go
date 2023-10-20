package suubdivision

import "dataon/division"

type Service interface {
	GetSubDivisions() ([]SubDivision, error)
	CreateSubDivision(input CreateSubDivisionInput, DivisionID int) (SubDivision, error)
	DeleteSubDivision(subDivisionID int) error
	UpdateSubDivision(inputID SubDivisionIDInput, inputData CreateSubDivisionInput) (SubDivision, error)
	GetSubDivisionsByDivisionID(divisionID int) ([]SubDivisionList, error)
}

type service struct {
	repository         Repository
	divisionRepository division.Repository
}

func NewService(repository Repository, divisionRepository division.Repository) *service {
	return &service{repository, divisionRepository}
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

func (s *service) GetSubDivisionsByDivisionID(divisionID int) ([]SubDivisionList, error) {
	division, err := s.divisionRepository.FindByID(divisionID)
	if err != nil {
		return nil, err
	}

	var subDivisionLists []SubDivisionList

	subDivisions, err := s.repository.FindAllByDivisionID(divisionID)
	if err != nil {
		return subDivisionLists, err
	}

	subList := SubDivisionList{
		DivisionID:           division.ID,
		ExecutiveCommitteeID: division.ExecutiveCommitteeID,
		DivisionName:         division.Name,
		SubDivision:          []SubDivision{},
	}

	for _, subDivision := range subDivisions {
		sub := SubDivision{
			ID:   subDivision.ID,
			Name: subDivision.Name,
		}

		subList.SubDivision = append(subList.SubDivision, sub)
	}

	subDivisionLists = append(subDivisionLists, subList)

	return subDivisionLists, nil
}
