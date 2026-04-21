package product_get

import (
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type InputProductGet struct {
	ProductID uuid.UUID `path:"id" json:"id" doc:"productID to get" format:"uuid"`
}

type OutputProductGet struct {
	Body struct {
		Product *domain.Product `json:"product" doc:"the requested product"`
	}
}
