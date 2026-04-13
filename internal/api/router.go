package api

import (
	"net/http"
	"shopapi/config"
	"shopapi/internal/core/client/client_create"
	"shopapi/internal/core/client/client_delete"
	"shopapi/internal/core/client/client_update_address"
	"shopapi/internal/core/client/clients_get_list"
	"shopapi/internal/core/client/clients_search"
	"shopapi/internal/core/image/image_assign"
	"shopapi/internal/core/image/image_delete"
	"shopapi/internal/core/image/image_get"
	"shopapi/internal/core/image/image_get_by_product"
	"shopapi/internal/core/image/image_update"
	"shopapi/internal/core/product/product_create"
	"shopapi/internal/core/product/product_delete"
	"shopapi/internal/core/product/product_get"
	"shopapi/internal/core/product/products_get_list"
	"shopapi/internal/core/supplier/supplier_create"
	"shopapi/internal/core/supplier/supplier_delete"
	"shopapi/internal/core/supplier/supplier_get"
	"shopapi/internal/core/supplier/supplier_update_address"
	"shopapi/internal/core/supplier/suppliers_get_list"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func GetRouter(app config.App) http.Handler {
	mux := chi.NewMux()
	config := huma.DefaultConfig("SHOP API", "1.0.0")
	config.DocsPath = "/swagger"
	config.Info = getAdditionalInfo(config.Info)
	api := humachi.New(mux, config)

	grp := huma.NewGroup(api, "/api")
	grp_v1 := huma.NewGroup(grp, "/v1")
	{
		grpClients := huma.NewGroup(grp_v1, "/clients")
		client_create.RegisterHTTPv1Handler(grpClients, "/", http.MethodPost, app.Store)
		client_delete.RegisterHTTPv1Handler(grpClients, "/{id}", http.MethodDelete, app.Store)
		clients_search.RegisterHTTPv1Handler(grpClients, "/search", http.MethodPost, app.Store)
		clients_get_list.RegisterHTTPv1Handler(grpClients, "/", http.MethodGet, app.Store)
		client_update_address.RegisterHTTPv1Handler(grpClients, "/{id}/address", http.MethodPut, app.Store)

		grpProducts := huma.NewGroup(grp_v1, "/products")
		product_create.RegisterHTTPv1Handler(grpProducts, "/", http.MethodPost, app.Store)
		product_get.RegisterHTTPv1Handler(grpProducts, "/{id}", http.MethodGet, app.Store)
		products_get_list.RegisterHTTPv1Handler(grpProducts, "/", http.MethodGet, app.Store)
		product_delete.RegisterHTTPv1Handler(grpProducts, "/{id}", http.MethodDelete, app.Store)
		image_assign.RegisterHTTPv1Handler(grpProducts, "/{id}/image", http.MethodPut, app.Store)
		image_get_by_product.RegisterHTTPv1Handler(grpProducts, "/{id}/image", http.MethodGet, app.Store)

		grpSuppliers := huma.NewGroup(grp_v1, "/suppliers")
		supplier_create.RegisterHTTPv1Handler(grpSuppliers, "/", http.MethodPost, app.Store)
		supplier_update_address.RegisterHTTPv1Handler(grpSuppliers, "/{id}/address", http.MethodPatch, app.Store)
		supplier_delete.RegisterHTTPv1Handler(grpSuppliers, "/{id}", http.MethodDelete, app.Store)
		suppliers_get_list.RegisterHTTPv1Handler(grpSuppliers, "/", http.MethodGet, app.Store)
		supplier_get.RegisterHTTPv1Handler(grpSuppliers, "/{id}", http.MethodGet, app.Store)

		grpImages := huma.NewGroup(grp_v1, "/images")
		image_update.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodPut, app.Store)
		image_delete.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodDelete, app.Store)
		image_get.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodGet, app.Store)
	}
	return mux
}
