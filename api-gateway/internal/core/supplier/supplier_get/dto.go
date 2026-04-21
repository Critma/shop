package supplier_get

import (
	"shopapi/internal/domain"
	"github.com/google/uuid"
)

type InputSupplierGet struct {
	SupplierID uuid.UUID `path:"id" json:"id" doc:"supplierID to get" format:"uuid"`
}

type OutputSupplierGet struct {
	Body struct {
		Supplier *domain.Supplier `json:"supplier" doc:"the requested supplier"`
	}
}