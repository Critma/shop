package product_stock_decrease

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterHTTPv1Handler(api huma.API, path, method string, store Store) {
	uc := newUseCase(store)

	huma.Register(api, huma.Operation{
		OperationID:   "decrease-stock-product_v1",
		Method:        method,
		Path:          path,
		Summary:       "Decrease product stock",
		Description:   "Decrease product stock",
		Tags:          []string{"Products"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *InputProductStockDecrease) (*OutputProductStockDecrease, error) {
		output, err := uc.ProductCreate(ctx, i)
		if err != nil {
			return nil, err
		}
		return output, nil
	})
}
