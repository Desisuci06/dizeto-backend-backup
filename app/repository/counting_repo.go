package repository

import (
	model "dizeto-backend/app/model/counting"

	"github.com/jinzhu/gorm"
)

type CountingRepository interface {
	CreateCounting(counting *model.Counting) error
	GetAllCounting() ([]*model.Counting, error)
	GetCountingByID(id string) (*model.Counting, error)
	UpdateCounting(counting *model.Counting) error
}

type countingRepository struct {
	db *gorm.DB
}

func NewCountingRepository(db *gorm.DB) CountingRepository {
	return &countingRepository{db: db}
}

func (cr *countingRepository) CreateCounting(counting *model.Counting) error {
	if err := counting.Validate(); err != nil {
		return err
	}
	err := cr.db.Create(counting).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *countingRepository) GetAllCounting() ([]*model.Counting, error) {
	var countings []*model.Counting
	if err := cr.db.Find(&countings).Error; err != nil {
		return nil, err
	}
	return countings, nil
}

func (cr *countingRepository) GetCountingByID(id string) (*model.Counting, error) {
	var counting model.Counting
	err := cr.db.Where("id = ?", id).First(&counting).Error
	return &counting, err
}

func (cr *countingRepository) UpdateCounting(counting *model.Counting) error {

	if err := counting.Validate(); err != nil {
		return err
	}

	err := cr.db.Save(counting).Error
	return err
}
