package product_delete

import (
	"github.com/google/uuid"
)

type InputProductDelete struct {
	ProductID uuid.UUID `path:"id" json:"id" doc:"productID to delete" format:"uuid"`
}