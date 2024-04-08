package repository

import (
	model "dizeto-backend/app/model/about"

	"github.com/jinzhu/gorm"
)

type AboutRepository interface {
	CreateAbout(about *model.About) error
	GetAllAbout() ([]*model.About, error)
	GetAboutByID(id string) (*model.About, error)
	UpdateAbout(about *model.About) error
}

type aboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return &aboutRepository{db: db}
}

func (ar *aboutRepository) CreateAbout(about *model.About) error {
	if err := about.Validate(); err != nil {
		return err
	}
	err := ar.db.Create(about).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *aboutRepository) GetAllAbout() ([]*model.About, error) {
	var about []*model.About
	err := ar.db.Find(&about).Error
	return about, err
}

func (ar *aboutRepository) GetAboutByID(id string) (*model.About, error) {
	var about model.About
	err := ar.db.Where("id = ?", id).First(&about).Error
	return &about, err
}

func (ar *aboutRepository) UpdateAbout(about *model.About) error {

	if err := about.Validate(); err != nil {
		return err
	}

	err := ar.db.Save(about).Error
	return err
}
