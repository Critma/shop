package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID             uuid.UUID  `json:"id"`
	Name           string     `json:"name"`
	Category       string     `json:"category"`
	Price          float64    `json:"price"`
	AvailableStock int        `json:"availableStock"`
	LastUpdateDate time.Time  `json:"lastUpdateDate"`
	SupplierID     uuid.UUID  `json:"supplierID"`
	Supplier       Supplier   `json:"supplier"`
	ImageID        *uuid.UUID `json:"imageID,omitempty"`
	Image          *Image     `json:"image,omitempty"`
}
