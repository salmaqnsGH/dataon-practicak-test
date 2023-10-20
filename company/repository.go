package company

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Company, error)
	Save(company Company) (Company, error)
	FindByID(ID int) (Company, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Company, error) {
	var companies []Company
	err := r.db.Find(&companies).Error

	if err != nil {
		return companies, err
	}

	return companies, nil
}

func (r *repository) Save(company Company) (Company, error) {
	err := r.db.Create(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil
}

func (r *repository) FindByID(ID int) (Company, error) {
	var company Company
	err := r.db.Where("id = ?", ID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil
}
