package supplier_create

import (
	"shopapi/internal/domain"

	"github.com/google/uuid"
)

type InputSupplierCreate struct {
	Body struct {
		Name string `json:"name" minLength:"2" maxLength:"100" doc:"Supplier name" example:"Gazprom"`
		// Country     string `json:"country" doc:"Full name country" maxLength:"100" example:"Russia"`
		// City        string `json:"city" doc:"Full name city" maxLength:"100" example:"Moscow"`
		// Street      string `json:"street" doc:"Full name street" maxLength:"255" example:"Lenin street"`
		PhoneNumber string `json:"phone_number" doc:"Supplier phone number" example:"+79991234567"`
		// AddressID   *uuid.UUID `json:"address_id,omitempty" format:"uuid" doc:"address ID"`
	}
}

type OutputSupplierCreate struct {
	Body struct {
		ID uuid.UUID `json:"id" format:"uuid" doc:"created supplier ID"`
	}
}

func toSupplier(input *InputSupplierCreate) *domain.Supplier {
	return &domain.Supplier{
		Name:        input.Body.Name,
		PhoneNumber: input.Body.PhoneNumber,
		// AddressID:   input.Body.AddressID,
	}
}
