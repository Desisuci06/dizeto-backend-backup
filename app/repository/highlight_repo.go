package repository

import (
	model "dizeto-backend/app/model/highlight_porto"

	"github.com/jinzhu/gorm"
)

type HighlightRepository interface {
	CreateHighlight(highlight *model.HighlightPortofolio) error
	GetAllHighlight() ([]*model.HighlightPortofolio, error)
	GetHighlightByID(id string) (*model.HighlightPortofolio, error)
	UpdateHighlight(highlight *model.HighlightPortofolio) error
}

type highlightRepository struct {
	db *gorm.DB
}

func NewHighlightPortofolio(db *gorm.DB) HighlightRepository {
	return &highlightRepository{db: db}
}

func (hr *highlightRepository) CreateHighlight(highlight *model.HighlightPortofolio) error {
	if err := highlight.Validate(); err != nil {
		return err
	}
	err := hr.db.Create(highlight).Error
	if err != nil {
		return err
	}
	return nil
}

func (hr *highlightRepository) GetAllHighlight() ([]*model.HighlightPortofolio, error) {
	var highlights []*model.HighlightPortofolio
	err := hr.db.Find(&highlights).Error
	return highlights, err
}

func (hr *highlightRepository) GetHighlightByID(id string) (*model.HighlightPortofolio, error) {
	var highlight model.HighlightPortofolio
	err := hr.db.Where("id = ?", id).First(&highlight).Error
	return &highlight, err
}

func (hr *highlightRepository) UpdateHighlight(highlight *model.HighlightPortofolio) error {

	if err := highlight.Validate(); err != nil {
		return err
	}

	err := hr.db.Save(highlight).Error
	return err
}
