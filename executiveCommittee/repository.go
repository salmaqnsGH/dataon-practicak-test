package executiveCommittee

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Executivecommittee, error)
	Save(executivecommittee Executivecommittee) (Executivecommittee, error)
	FindAllByCompanyID(companyID int) ([]Executivecommittee, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Executivecommittee, error) {
	var Executivecommitties []Executivecommittee
	err := r.db.Find(&Executivecommitties).Error

	if err != nil {
		return Executivecommitties, err
	}

	return Executivecommitties, nil
}

func (r *repository) Save(executivecommittee Executivecommittee) (Executivecommittee, error) {
	err := r.db.Create(&executivecommittee).Error

	if err != nil {
		return executivecommittee, err
	}

	return executivecommittee, nil
}

func (r *repository) FindAllByCompanyID(companyID int) ([]Executivecommittee, error) {
	var Executivecommitties []Executivecommittee
	err := r.db.Where("company_id = ?", companyID).Find(&Executivecommitties).Error

	if err != nil {
		return Executivecommitties, err
	}

	return Executivecommitties, nil
}
