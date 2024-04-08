package repository

import (
	model "dizeto-backend/app/model/testimoni"

	"github.com/jinzhu/gorm"
)

type TestimoniRepository interface {
	CreateTestimoni(testimoni *model.Testimoni) error
	GetAllTestimoni() ([]*model.Testimoni, error)
	GetTestimoniByID(id string) (*model.Testimoni, error)
	UpdateTestimoni(testimoni *model.Testimoni) error
}

type testimoniRepository struct {
	db *gorm.DB
}

func NewTestimoniRepository(db *gorm.DB) TestimoniRepository {
	return &testimoniRepository{db: db}
}

func (tr *testimoniRepository) CreateTestimoni(testimoni *model.Testimoni) error {
	if err := testimoni.Validate(); err != nil {
		return err
	}
	err := tr.db.Create(testimoni).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *testimoniRepository) GetAllTestimoni() ([]*model.Testimoni, error) {
	var testimonis []*model.Testimoni
	if err := tr.db.Find(&testimonis).Error; err != nil {
		return nil, err
	}
	return testimonis, nil
}

func (tr *testimoniRepository) GetTestimoniByID(id string) (*model.Testimoni, error) {
	var testimoni model.Testimoni
	err := tr.db.Where("id = ?", id).First(&testimoni).Error
	return &testimoni, err
}

func (tr *testimoniRepository) UpdateTestimoni(testimoni *model.Testimoni) error {

	if err := testimoni.Validate(); err != nil {
		return err
	}

	err := tr.db.Save(testimoni).Error
	return err
}
