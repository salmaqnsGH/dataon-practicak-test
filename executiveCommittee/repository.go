package executiveCommittee

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Executivecommittee, error)
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
