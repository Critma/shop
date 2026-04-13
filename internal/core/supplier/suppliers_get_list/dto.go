package suppliers_get_list

import (
	"shopapi/internal/domain"
)

type InputGetAllSuppliers struct {
	Limit  int64 `query:"limit,omitzero" doc:"number of suppliers to return"`
	Offset int64 `query:"offset,omitzero" doc:"number of suppliers to skip"`
}

type OutputGetAllSuppliers struct {
	Body struct {
		Suppliers []*domain.Supplier `json:"suppliers" doc:"list of all suppliers"`
	}
}