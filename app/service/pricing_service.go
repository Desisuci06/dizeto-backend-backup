package service

import (
	model "dizeto-backend/app/model/pricing"
	"dizeto-backend/app/model/pricing/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type PricingService interface {
	CreatePricing(title, paket, category, item_list string, price uint) error
	GetAllPricing() (*dto.ResponsePricingsDTO, error)
	UpdatePricing(id, title, paket, category, item_list string, price uint) error
}

type pricingService struct {
	pricingRepo repository.PricingRepository
}

func NewPricingService(pricingRepo repository.PricingRepository) PricingService {
	return &pricingService{pricingRepo: pricingRepo}
}

func (ps *pricingService) CreatePricing(title, paket, category, item_list string, price uint) error {
	// Generate UUID for pricing ID
	pricingID := uuid.New()

	// Create new pricing
	newPricing := &model.Pricing{
		ID:       pricingID,
		Title:    title,
		Price:    price,
		Paket:    paket,
		Category: category,
		ItemList: item_list,
		PageID:   1,
	}

	// Save new pricing to repository
	err := ps.pricingRepo.CreatePricing(newPricing)
	if err != nil {
		return err
	}

	return nil
}

func (ps *pricingService) GetAllPricing() (*dto.ResponsePricingsDTO, error) {
	pricings, err := ps.pricingRepo.GetAllPricing()
	if err != nil {
		return nil, err
	}

	var pricingDTOs []*dto.ResponseDTO
	for _, p := range pricings {
		pricingDTO := &dto.ResponseDTO{
			ID:       p.ID,
			Title:    p.Title,
			Price:    p.Price,
			Paket:    p.Paket,
			Category: p.Category,
			ItemList: p.ItemList,
		}
		pricingDTOs = append(pricingDTOs, pricingDTO)
	}

	responseDTO := &dto.ResponsePricingsDTO{
		Pricings: pricingDTOs,
	}
	return responseDTO, nil
}

func (ps *pricingService) UpdatePricing(id, title, paket, category, item_list string, price uint) error {
	pricing, err := ps.pricingRepo.GetPricingByID(id)
	if err != nil {
		return err
	}

	pricing.Title = title
	pricing.Price = price
	pricing.Paket = paket
	pricing.Category = category
	pricing.ItemList = item_list

	if err := ps.pricingRepo.UpdatePricing(pricing); err != nil {
		return err
	}

	return nil
}
