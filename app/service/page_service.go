package service

import (
	model "dizeto-backend/app/model/page"
	"dizeto-backend/app/repository"
)

type LandingService interface {
	GetLandingPage() ([]*model.Page, error)
}

type landingService struct {
	landingRepo repository.LandingRepository
}

func NewLandingService(landingRepo repository.LandingRepository) LandingService {
	return &landingService{landingRepo: landingRepo}
}

func (ls *landingService) GetLandingPage() ([]*model.Page, error) {
	return ls.landingRepo.GetLandingPage()
}
