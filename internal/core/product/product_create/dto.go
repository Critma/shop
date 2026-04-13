package product_create

import (
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type InputProductCreate struct {
	Body struct {
		Name           string    `json:"name" minLength:"2" maxLength:"100" doc:"Product name" example:"Dishwasher"`
		Category       string    `json:"category" minLength:"2" maxLength:"100" doc:"Product category" example:"Electronics"`
		Price          float64   `json:"price" doc:"Product price" minimum:"0"`
		AvailableStock int       `json:"availableStock" doc:"Product available stock" minimum:"0" example:"100"`
		SupplierID     uuid.UUID `json:"supplierID" format:"uuid" doc:"supplier ID"`
		// ImageID        *uuid.UUID `json:"imageID,omitempty" format:"uuid"`
	}
}

type OutputProductCreate struct {
	Body struct {
		ID uuid.UUID `json:"id" format:"uuid" doc:"product ID"`
	}
}

func toProduct(input *InputProductCreate) *domain.Product {
	return &domain.Product{
		Name:           input.Body.Name,
		Category:       input.Body.Category,
		Price:          input.Body.Price,
		AvailableStock: input.Body.AvailableStock,
		SupplierID:     input.Body.SupplierID,
		// ImageID:        input.Body.ImageID,
	}
}
