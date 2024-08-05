package service

import (
	model "dizeto-backend/app/model/client"
	"dizeto-backend/app/model/client/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type ClientService interface {
	CreateClient(title, logoURL, theme, href, alt string) error
	GetAllClient() (*dto.ResponseClientsDTO, error)
	UpdateClient(id, title, logoURL, theme, href, alt string) error
	GetClientByID(id string) (*model.Client, error)
}

type clienttService struct {
	clienttRepo repository.ClientRepository
	titleRepo   repository.TitleRepository
}

func NewClientService(clienttRepo repository.ClientRepository, titleRepo repository.TitleRepository) ClientService {
	return &clienttService{
		clienttRepo: clienttRepo,
		titleRepo:   titleRepo,
	}
}

func (hs *clienttService) CreateClient(title, logoURL, theme, href, alt string) error {
	// Generate UUID for clientt ID
	clienttID := uuid.New()

	// Create new clientt
	newClient := &model.Client{
		ID:      clienttID,
		Title:   title,
		LogoURL: logoURL,
		Theme:   theme,
		Href:    href,
		Alt:     alt,
		PageID:  1,
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
		title, err := hs.titleRepo.GetTitleByKdTitle(p.Title)
		if err != nil {
			return nil, err
		}
		clienttDTO := &dto.ResponseDTO{
			ID:      p.ID,
			Title:   title.NmTitle,
			LogoURL: p.LogoURL,
			Theme:   p.Theme,
			Href:    p.Href,
			Alt:     p.Alt,
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

func (hs *clienttService) UpdateClient(id, title, logoURL, theme, href, alt string) error {
	clientt, err := hs.clienttRepo.GetClientByID(id)
	if err != nil {
		return err
	}

	clientt.Title = title
	clientt.LogoURL = logoURL
	clientt.Theme = theme
	clientt.Href = href
	clientt.Alt = alt

	if err := hs.clienttRepo.UpdateClient(clientt); err != nil {
		return err
	}

	return nil
}
