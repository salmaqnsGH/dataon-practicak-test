package suubdivision

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]SubDivision, error)
	Save(subDivision SubDivision) (SubDivision, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]SubDivision, error) {
	var subDivisions []SubDivision
	err := r.db.Find(&subDivisions).Error

	if err != nil {
		return subDivisions, err
	}

	return subDivisions, nil
}

func (r *repository) Save(subDivision SubDivision) (SubDivision, error) {
	err := r.db.Create(&subDivision).Error

	if err != nil {
		return subDivision, err
	}

	return subDivision, nil
}
