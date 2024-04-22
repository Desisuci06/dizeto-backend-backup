package service

import (
	model "dizeto-backend/app/model/title"
	"dizeto-backend/app/model/title/dto"

	"dizeto-backend/app/repository"
)

type TitleService interface {
	CreateTitle(id int, KdTitle, NmTitle string) error
	GetAllTitle() (*dto.ResponseTitlesDTO, error)
	UpdateTitle(id int, KdTitle, NmTitle string) error
}

type titleService struct {
	titleRepo repository.TitleRepository
}

func NewTitleService(titleRepo repository.TitleRepository) TitleService {
	return &titleService{titleRepo: titleRepo}
}

func (as *titleService) CreateTitle(id int, KdTitle, NmTitle string) error {

	// Create new title
	newTitle := &model.Title{
		ID:      id,
		KdTitle: KdTitle,
		NmTitle: NmTitle,
	}

	// Save new title to repository
	err := as.titleRepo.CreateTitle(newTitle)
	if err != nil {
		return err
	}

	return nil
}

func (as *titleService) GetAllTitle() (*dto.ResponseTitlesDTO, error) {
	titles, err := as.titleRepo.GetAllTitle()
	if err != nil {
		return nil, err
	}

	var titleDTOs []*dto.ResponseDTO
	for _, p := range titles {
		titleDTO := &dto.ResponseDTO{
			ID:      p.ID,
			KdTitle: p.KdTitle,
			NmTitle: p.NmTitle,
		}
		titleDTOs = append(titleDTOs, titleDTO)
	}

	responseDTO := &dto.ResponseTitlesDTO{
		Titles: titleDTOs,
	}
	return responseDTO, nil
}

func (as *titleService) UpdateTitle(id int, KdTitle, NmTitle string) error {
	title, err := as.titleRepo.GetTitleByID(id)
	if err != nil {
		return err
	}

	title.KdTitle = KdTitle
	title.NmTitle = NmTitle

	if err := as.titleRepo.UpdateTitle(title); err != nil {
		return err
	}

	return nil
}
