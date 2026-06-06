package product_stock_decrease

import (
	"github.com/google/uuid"
)

type InputProductStockDecrease struct {
	Body struct {
		ProductID uuid.UUID `json:"productID" format:"uuid" doc:"product ID"`
		Quantity  int       `json:"quantity" doc:"quantity to decrease" minimum:"0"`
	}
}

type OutputProductStockDecrease struct {
	Body struct {
	}
}
