package products_get_list

import (
	"shopapi/internal/domain"
)

type InputGetAllProducts struct {
	Limit  int64 `query:"limit,omitzero" doc:"number of products to return"`
	Offset int64 `query:"offset,omitzero" doc:"number of products to skip"`
}

type OutputGetAllProducts struct {
	Body struct {
		Products []*domain.Product `json:"products" doc:"list of all products"`
	}
}