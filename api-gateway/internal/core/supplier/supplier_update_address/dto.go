package supplier_update_address

import (
	"shopapi/internal/domain"
	"github.com/google/uuid"
)

type InputUpdateSupplierAddress struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"supplier ID"`
	Body struct {
		Country string `json:"country" doc:"Full name country" maxLength:"100" example:"Russia"`
		City    string `json:"city" doc:"Full name city" maxLength:"100" example:"Moscow"`
		Street  string `json:"street" doc:"Full name street" maxLength:"255" example:"Walter street"`
	} `doc:"new address for the supplier"`
}

type OutputUpdateSupplierAddress struct {
	Body struct {
		Success bool `json:"success" doc:"indicates if the update was successful"`
	}
}

func toAddress(input *InputUpdateSupplierAddress) *domain.Address {
	return &domain.Address{
		Country: input.Body.Country,
		City:    input.Body.City,
		Street:  input.Body.Street,
	}
}