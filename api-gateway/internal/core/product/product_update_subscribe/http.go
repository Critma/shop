package product_update_subscribe

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/sse"
)

// func RegisterHTTPv1Handler(api huma.API, path, method string) {
// 	huma.Register(api, huma.Operation{
// 		OperationID:   "ws-product-update-broadcast_v1",
// 		Method:        method,
// 		Path:          path,
// 		Summary:       "Get all products",
// 		Description:   "Get all products with optional pagination (limit and offset)",
// 		Tags:          []string{"Products"},
// 		DefaultStatus: http.StatusOK,
// 		Security: []map[string][]string{
// 			{"bearer": {}},
// 		},
// 	}, func(ctx context.Context, i *InputGetAllProducts) (*OutputGetAllProducts, error) {
// 		output, err := usecase.GetAllProducts(ctx, i)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return output, nil
// 	})
// }

func RegisterSSEHandler(api huma.API, path string) {
	sse.Register(api, huma.Operation{
		OperationID: "sse-product-update-broadcast",
		Method:      http.MethodGet,
		Path:        path,
		Summary:     "Subscribe to product updates",
		Tags:        []string{"Products"},
	}, map[string]any{
		"message": SEEMessage{},
	}, func(ctx context.Context, input *struct{}, send sse.Sender) {
		productUpdateChan := usecase.ListenSSEBroadcast(context.Background())
		for productUpdate := range productUpdateChan {
			send.Data(productUpdate)
		}
	})
}
