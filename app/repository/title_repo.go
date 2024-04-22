package repository

import (
	model "dizeto-backend/app/model/title"

	"github.com/jinzhu/gorm"
)

type TitleRepository interface {
	CreateTitle(title *model.Title) error
	GetAllTitle() ([]*model.Title, error)
	GetTitleByID(id int) (*model.Title, error)
	UpdateTitle(title *model.Title) error
	GetTitleByKdTitle(kd_title string) (*model.Title, error)
}

type titleRepository struct {
	db *gorm.DB
}

func NewTitleRepository(db *gorm.DB) TitleRepository {
	return &titleRepository{db: db}
}

func (ar *titleRepository) CreateTitle(title *model.Title) error {
	if err := title.Validate(); err != nil {
		return err
	}
	err := ar.db.Create(title).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *titleRepository) GetAllTitle() ([]*model.Title, error) {
	var title []*model.Title
	err := ar.db.Find(&title).Error
	return title, err
}

func (ar *titleRepository) GetTitleByID(id int) (*model.Title, error) {
	var title model.Title
	err := ar.db.Where("id = ?", id).First(&title).Error
	return &title, err
}

func (ar *titleRepository) GetTitleByKdTitle(kd_title string) (*model.Title, error) {
	var title model.Title
	err := ar.db.Where("kd_title = ?", kd_title).First(&title).Error
	return &title, err
}

func (ar *titleRepository) UpdateTitle(title *model.Title) error {

	if err := title.Validate(); err != nil {
		return err
	}

	err := ar.db.Save(title).Error
	return err
}
