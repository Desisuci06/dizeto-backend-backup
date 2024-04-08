package service

import (
	model "dizeto-backend/app/model/counting"
	"dizeto-backend/app/model/counting/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type CountingService interface {
	CreateCounting(happy_client, completed_projects, subscribers uint) error
	GetAllCounting() (*dto.ResponseCountingsDTO, error)
	UpdateCounting(id string, happy_client, completed_projects, subscribers uint) error
}

type countingService struct {
	countingRepo repository.CountingRepository
}

func NewCountingService(countingRepo repository.CountingRepository) CountingService {
	return &countingService{countingRepo: countingRepo}
}

func (cs *countingService) CreateCounting(happy_client, completed_projects, subscribers uint) error {
	// Generate UUID for counting ID
	countingID := uuid.New()

	// Create new counting
	newCounting := &model.Counting{
		ID:                countingID,
		HappyClient:       happy_client,
		CompletedProjects: completed_projects,
		Subscribers:       subscribers,
		PageID:            1,
	}

	// Save new counting to repository
	err := cs.countingRepo.CreateCounting(newCounting)
	if err != nil {
		return err
	}

	return nil
}

func (cs *countingService) GetAllCounting() (*dto.ResponseCountingsDTO, error) {
	countings, err := cs.countingRepo.GetAllCounting()
	if err != nil {
		return nil, err
	}

	var countingDTOs []*dto.ResponseDTO
	for _, p := range countings {
		countingDTO := &dto.ResponseDTO{
			ID:                p.ID,
			HappyClient:       p.HappyClient,
			CompletedProjects: p.CompletedProjects,
			Subscribers:       p.Subscribers,
		}
		countingDTOs = append(countingDTOs, countingDTO)
	}

	responseDTO := &dto.ResponseCountingsDTO{
		Countings: countingDTOs,
	}
	return responseDTO, nil
}

func (cs *countingService) UpdateCounting(id string, happy_client, completed_projects, subscribers uint) error {
	counting, err := cs.countingRepo.GetCountingByID(id)
	if err != nil {
		return err
	}

	counting.HappyClient = happy_client
	counting.CompletedProjects = completed_projects
	counting.Subscribers = subscribers

	if err := cs.countingRepo.UpdateCounting(counting); err != nil {
		return err
	}

	return nil
}
