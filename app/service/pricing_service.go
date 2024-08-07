package service

import (
	model_item_list "dizeto-backend/app/model/item"
	model "dizeto-backend/app/model/pricing"
	"dizeto-backend/app/model/pricing/dto"
	"fmt"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type PricingService interface {
	CreatePricing(pricingID uuid.UUID, title, paket, category string, itemList []*model_item_list.ItemList, price uint) error
	GetAllPricing() (*dto.ResponsePricingsDTO, error)
	UpdatePricing(id string, itemList []*model_item_list.ItemList, pricingDTO *dto.PricingDTO) error
}

type pricingService struct {
	pricingRepo repository.PricingRepository
	titleRepo   repository.TitleRepository
}

func NewPricingService(pricingRepo repository.PricingRepository, titleRepo repository.TitleRepository) PricingService {
	return &pricingService{
		pricingRepo: pricingRepo,
		titleRepo:   titleRepo,
	}
}

func (ps *pricingService) CreatePricing(pricingID uuid.UUID, title, paket, category string, itemList []*model_item_list.ItemList, price uint) error {
	// Create new pricing
	newPricing := &model.Pricing{
		ID:       pricingID,
		Title:    title,
		Price:    price,
		Paket:    paket,
		Category: category,
		ItemList: itemList,
		PageID:   1,
	}

	// Save new pricing to repository
	err := ps.pricingRepo.CreatePricing(newPricing, itemList)
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
		title, err := ps.titleRepo.GetTitleByKdTitle(p.Title)
		if err != nil {
			return nil, err
		}
		pricingDTO := &dto.ResponseDTO{
			ID:       p.ID,
			Title:    title.NmTitle,
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

func (ps *pricingService) UpdatePricing(id string, itemList []*model_item_list.ItemList, pricingDTO *dto.PricingDTO) error {
	// Retrieve pricing by ID
	pricing, err := ps.pricingRepo.GetPricingByID(id)
	if err != nil {
		return err
	}

	// Update pricing fields
	pricing.Price = pricingDTO.Price
	pricing.Paket = pricingDTO.Paket
	pricing.Category = pricingDTO.Category

	// Validate and update item list
	updatedItems := make([]*model_item_list.ItemList, 0)
	for _, newItem := range itemList {
		found := false
		for _, existingItem := range pricing.ItemList {
			if newItem.ID == existingItem.ID {
				found = true
				updatedItems = append(updatedItems, newItem)
				break
			}
		}
		if !found {
			return fmt.Errorf("item with ID %s not found", newItem.ID)
		}
	}

	// Replace the existing item list with the updated list
	pricing.ItemList = updatedItems

	// Validate pricing entity
	if err := pricing.Validate(); err != nil {
		return err
	}

	// Update pricing in repository
	err = ps.pricingRepo.UpdatePricing(pricing)
	if err != nil {
		return err
	}

	return nil
}
