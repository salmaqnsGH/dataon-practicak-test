package suubdivision

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]SubDivision, error)
	Save(subDivision SubDivision) (SubDivision, error)
	Delete(subDivisionID int) error
	FindByID(ID int) (SubDivision, error)
	Update(subDivision SubDivision) (SubDivision, error)
	FindAllByDivisionID(divisionID int) ([]SubDivision, error)
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

func (r *repository) Delete(subDivisionID int) error {
	var subDivision SubDivision
	if err := r.db.Where("id = ?", subDivisionID).First(&subDivision).Delete(&subDivision).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(ID int) (SubDivision, error) {
	var subDivision SubDivision
	err := r.db.Where("id = ?", ID).Find(&subDivision).Error

	if err != nil {
		return subDivision, err
	}

	return subDivision, nil
}

func (r *repository) Update(subDivision SubDivision) (SubDivision, error) {
	err := r.db.Save(&subDivision).Error

	if err != nil {
		return subDivision, err
	}

	return subDivision, nil
}

func (r *repository) FindAllByDivisionID(divisionID int) ([]SubDivision, error) {
	var subDivisions []SubDivision
	err := r.db.Where("division_id = ?", divisionID).Find(&subDivisions).Error

	fmt.Println(subDivisions)

	if err != nil {
		return subDivisions, err
	}

	return subDivisions, nil
}
