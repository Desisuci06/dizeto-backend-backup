package service

import (
	model "dizeto-backend/app/model/highlight_porto"
	"dizeto-backend/app/model/highlight_porto/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type HighlightService interface {
	CreateHighlight(title, image string) error
	GetAllHighlight() (*dto.ResponseHighlightsDTO, error)
	UpdateHighlight(id, title, image string) error
	GetHighlightByID(id string) (*model.HighlightPortofolio, error)
}

type highlightService struct {
	highlightRepo repository.HighlightRepository
	titleRepo     repository.TitleRepository
}

func NewHighlightService(highlightRepo repository.HighlightRepository, titleRepo repository.TitleRepository) HighlightService {
	return &highlightService{
		highlightRepo: highlightRepo,
		titleRepo:     titleRepo,
	}
}

func (hs *highlightService) CreateHighlight(title, image string) error {
	// Generate UUID for highlight ID
	highlightID := uuid.New()

	// Create new highlight
	newHighlight := &model.HighlightPortofolio{
		ID:     highlightID,
		Title:  title,
		Image:  image,
		PageID: 1,
	}

	// Save new highlight to repository
	err := hs.highlightRepo.CreateHighlight(newHighlight)
	if err != nil {
		return err
	}

	return nil
}

func (hs *highlightService) GetAllHighlight() (*dto.ResponseHighlightsDTO, error) {
	highlights, err := hs.highlightRepo.GetAllHighlight()
	if err != nil {
		return nil, err
	}

	var highlightDTOs []*dto.ResponseDTO
	for _, p := range highlights {
		title, err := hs.titleRepo.GetTitleByKdTitle(p.Title)
		if err != nil {
			return nil, err
		}
		highlightDTO := &dto.ResponseDTO{
			ID:    p.ID,
			Title: title.NmTitle,
			Image: p.Image,
		}
		highlightDTOs = append(highlightDTOs, highlightDTO)
	}

	responseDTO := &dto.ResponseHighlightsDTO{
		Highlight_portofolio: highlightDTOs,
	}
	return responseDTO, nil
}

func (hs *highlightService) GetHighlightByID(id string) (*model.HighlightPortofolio, error) {
	highlight, err := hs.highlightRepo.GetHighlightByID(id)
	if err != nil {
		return nil, err
	}

	return highlight, nil
}

func (hs *highlightService) UpdateHighlight(id, title, image string) error {
	highlight, err := hs.highlightRepo.GetHighlightByID(id)
	if err != nil {
		return err
	}

	highlight.Title = title
	highlight.Image = image

	if err := hs.highlightRepo.UpdateHighlight(highlight); err != nil {
		return err
	}

	return nil
}
