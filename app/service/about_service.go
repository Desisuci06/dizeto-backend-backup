package service

import (
	model "dizeto-backend/app/model/about"
	"dizeto-backend/app/model/about/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type AboutService interface {
	CreateAbout(title, subtitle, description, note, image string) error
	GetAllAbout() (*dto.ResponseAboutsDTO, error)
	UpdateAbout(id, title, subtitle, description, note, image string) error
}

type aboutService struct {
	aboutRepo repository.AboutRepository
}

func NewAboutService(aboutRepo repository.AboutRepository) AboutService {
	return &aboutService{aboutRepo: aboutRepo}
}

func (as *aboutService) CreateAbout(title, subtitle, description, note, image string) error {
	// Generate UUID for about ID
	aboutID := uuid.New()

	// Create new about
	newAbout := &model.About{
		ID:          aboutID,
		Title:       title,
		Subtitle:    subtitle,
		Description: description,
		Note:        note,
		Image:       image,
		PageID:      1,
	}

	// Save new about to repository
	err := as.aboutRepo.CreateAbout(newAbout)
	if err != nil {
		return err
	}

	return nil
}

func (as *aboutService) GetAllAbout() (*dto.ResponseAboutsDTO, error) {
	abouts, err := as.aboutRepo.GetAllAbout()
	if err != nil {
		return nil, err
	}

	var aboutDTOs []*dto.ResponseDTO
	for _, p := range abouts {
		aboutDTO := &dto.ResponseDTO{
			ID:          p.ID,
			Title:       p.Title,
			Subtitle:    p.Subtitle,
			Description: p.Description,
			Note:        p.Note,
			Image:       p.Image,
		}
		aboutDTOs = append(aboutDTOs, aboutDTO)
	}

	responseDTO := &dto.ResponseAboutsDTO{
		Abouts: aboutDTOs,
	}
	return responseDTO, nil
}

func (as *aboutService) UpdateAbout(id, title, subtitle, description, note, image string) error {
	about, err := as.aboutRepo.GetAboutByID(id)
	if err != nil {
		return err
	}

	about.Title = title
	about.Subtitle = subtitle
	about.Description = description
	about.Note = note
	about.Image = image

	if err := as.aboutRepo.UpdateAbout(about); err != nil {
		return err
	}

	return nil
}
