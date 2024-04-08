package service

import (
	model "dizeto-backend/app/model/testimoni"
	"dizeto-backend/app/model/testimoni/dto"

	"dizeto-backend/app/repository"

	"github.com/google/uuid"
)

type TestimoniService interface {
	CreateTestimoni(image, name, event, comment string) error
	GetAllTestimoni() (*dto.ResponseTestimonisDTO, error)
	UpdateTestimoni(id, image, name, event, comment string) error
}

type testimoniService struct {
	testimoniRepo repository.TestimoniRepository
}

func NewTestimoniService(testimoniRepo repository.TestimoniRepository) TestimoniService {
	return &testimoniService{testimoniRepo: testimoniRepo}
}

func (ts *testimoniService) CreateTestimoni(image, name, event, comment string) error {
	// Generate UUID for testimoni ID
	testimoniID := uuid.New()

	// Create new testimoni
	newTestimoni := &model.Testimoni{
		ID:      testimoniID,
		Image:   image,
		Name:    name,
		Event:   event,
		Comment: comment,
		PageID:  1,
	}

	// Save new testimoni to repository
	err := ts.testimoniRepo.CreateTestimoni(newTestimoni)
	if err != nil {
		return err
	}

	return nil
}

func (ts *testimoniService) GetAllTestimoni() (*dto.ResponseTestimonisDTO, error) {
	testimonis, err := ts.testimoniRepo.GetAllTestimoni()
	if err != nil {
		return nil, err
	}

	var testimoniDTOs []*dto.ResponseDTO
	for _, p := range testimonis {
		testimoniDTO := &dto.ResponseDTO{
			ID:      p.ID,
			Image:   p.Image,
			Name:    p.Name,
			Event:   p.Event,
			Comment: p.Comment,
		}
		testimoniDTOs = append(testimoniDTOs, testimoniDTO)
	}

	responseDTO := &dto.ResponseTestimonisDTO{
		Testimonis: testimoniDTOs,
	}
	return responseDTO, nil
}

func (ts *testimoniService) UpdateTestimoni(id, image, name, event, comment string) error {
	testimoni, err := ts.testimoniRepo.GetTestimoniByID(id)
	if err != nil {
		return err
	}

	testimoni.Image = image
	testimoni.Name = name
	testimoni.Event = event
	testimoni.Comment = comment

	if err := ts.testimoniRepo.UpdateTestimoni(testimoni); err != nil {
		return err
	}

	return nil
}
