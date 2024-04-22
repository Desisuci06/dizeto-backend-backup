package repository

import (
	model "dizeto-backend/app/model/client"

	"github.com/jinzhu/gorm"
)

type ClientRepository interface {
	CreateClient(client *model.Client) error
	GetAllClient() ([]*model.Client, error)
	GetClientByID(id string) (*model.Client, error)
	UpdateClient(client *model.Client) error
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (cr *clientRepository) CreateClient(client *model.Client) error {
	if err := client.Validate(); err != nil {
		return err
	}
	err := cr.db.Create(client).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *clientRepository) GetAllClient() ([]*model.Client, error) {
	var client []*model.Client
	err := cr.db.
		Joins("JOIN titles ON clients.title = titles.kd_title").
		Select("clients.*, titles.nm_title as Title").
		Find(&client).Error

	return client, err
}

func (cr *clientRepository) GetClientByID(id string) (*model.Client, error) {
	var client model.Client
	err := cr.db.Where("id = ?", id).First(&client).Error
	return &client, err
}

func (cr *clientRepository) UpdateClient(client *model.Client) error {

	if err := client.Validate(); err != nil {
		return err
	}

	err := cr.db.Save(client).Error
	return err
}
