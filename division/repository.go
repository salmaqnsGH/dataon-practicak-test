package division

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Division, error)
	Save(division Division) (Division, error)
	FindByID(ID int) (Division, error)
	FindByECID(ID int) ([]Division, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Division, error) {
	var divisions []Division
	err := r.db.Find(&divisions).Error

	if err != nil {
		return divisions, err
	}

	return divisions, nil
}

func (r *repository) Save(division Division) (Division, error) {
	err := r.db.Create(&division).Error

	if err != nil {
		return division, err
	}

	return division, nil
}

func (r *repository) FindByID(ID int) (Division, error) {
	var division Division
	err := r.db.Where("id = ?", ID).Find(&division).Error

	if err != nil {
		return division, err
	}

	return division, nil
}

func (r *repository) FindByECID(ID int) ([]Division, error) {
	var division []Division
	err := r.db.Where("executive_committee_id = ?", ID).Find(&division).Error

	if err != nil {
		return division, err
	}

	return division, nil
}
