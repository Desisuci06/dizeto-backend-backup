package service

import (
	model "dizeto-backend/app/model/client"
	"dizeto-backend/app/model/client/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type ClientService interface {
	CreateClient(title, image string) error
	GetAllClient() (*dto.ResponseClientsDTO, error)
	UpdateClient(id, title, image string) error
	GetClientByID(id string) (*model.Client, error)
}

type clienttService struct {
	clienttRepo repository.ClientRepository
}

func NewClientService(clienttRepo repository.ClientRepository) ClientService {
	return &clienttService{clienttRepo: clienttRepo}
}

func (hs *clienttService) CreateClient(title, image string) error {
	// Generate UUID for clientt ID
	clienttID := uuid.New()

	// Create new clientt
	newClient := &model.Client{
		ID:     clienttID,
		Title:  title,
		Image:  image,
		PageID: 1,
	}

	// Save new clientt to repository
	err := hs.clienttRepo.CreateClient(newClient)
	if err != nil {
		return err
	}

	return nil
}

func (hs *clienttService) GetAllClient() (*dto.ResponseClientsDTO, error) {
	clientts, err := hs.clienttRepo.GetAllClient()
	if err != nil {
		return nil, err
	}

	var clienttDTOs []*dto.ResponseDTO
	for _, p := range clientts {
		clienttDTO := &dto.ResponseDTO{
			ID:    p.ID,
			Title: p.Title,
			Image: p.Image,
		}
		clienttDTOs = append(clienttDTOs, clienttDTO)
	}

	responseDTO := &dto.ResponseClientsDTO{
		Clients: clienttDTOs,
	}
	return responseDTO, nil
}

func (hs *clienttService) GetClientByID(id string) (*model.Client, error) {
	clientt, err := hs.clienttRepo.GetClientByID(id)
	if err != nil {
		return nil, err
	}

	return clientt, nil
}

func (hs *clienttService) UpdateClient(id, title, image string) error {
	clientt, err := hs.clienttRepo.GetClientByID(id)
	if err != nil {
		return err
	}

	clientt.Title = title
	clientt.Image = image

	if err := hs.clienttRepo.UpdateClient(clientt); err != nil {
		return err
	}

	return nil
}
