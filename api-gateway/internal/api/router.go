package api

import (
	"net/http"
	"shopapi/config"
	"shopapi/internal/core/auth/auth_login"
	"shopapi/internal/core/auth/auth_register"
	"shopapi/internal/core/auth/auth_reset"
	"shopapi/internal/core/auth/middleware/auth_header_middleware"
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

// TODO: refactor: move usecases out
func GetRouter(app config.App) http.Handler {
	mux := chi.NewMux()
	config := huma.DefaultConfig("SHOP API", "1.0.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	config.DocsPath = "/swagger"
	config.Info = getAdditionalInfo(config.Info)

	api := humachi.New(mux, config)

	authMiddleware := auth_header_middleware.AuthMiddleware(api)

	grp := huma.NewGroup(api, "/api")
	grp_v1 := huma.NewGroup(grp, "/v1")
	{
		grpAuthNeed := huma.NewGroup(grp_v1)
		grpAuthNeed.UseMiddleware(authMiddleware)

		grpClients := huma.NewGroup(grpAuthNeed, "/clients")
		client_create.RegisterHTTPv1Handler(grpClients, "/", http.MethodPost)
		client_delete.RegisterHTTPv1Handler(grpClients, "/{id}", http.MethodDelete)
		clients_search.RegisterHTTPv1Handler(grpClients, "/search", http.MethodPost)
		clients_get_list.RegisterHTTPv1Handler(grpClients, "/", http.MethodGet)
		client_update_address.RegisterHTTPv1Handler(grpClients, "/{id}/address", http.MethodPut)

		grpProducts := huma.NewGroup(grpAuthNeed, "/products")
		product_create.RegisterHTTPv1Handler(grpProducts, "/", http.MethodPost)
		product_get.RegisterHTTPv1Handler(grpProducts, "/{id}", http.MethodGet)
		products_get_list.RegisterHTTPv1Handler(grpProducts, "/", http.MethodGet)
		product_delete.RegisterHTTPv1Handler(grpProducts, "/{id}", http.MethodDelete)
		image_assign.RegisterHTTPv1Handler(grpProducts, "/{id}/image", http.MethodPut)
		image_get_by_product.RegisterHTTPv1Handler(grpProducts, "/{id}/image", http.MethodGet)

		grpSuppliers := huma.NewGroup(grpAuthNeed, "/suppliers")
		supplier_create.RegisterHTTPv1Handler(grpSuppliers, "/", http.MethodPost)
		supplier_update_address.RegisterHTTPv1Handler(grpSuppliers, "/{id}/address", http.MethodPatch)
		supplier_delete.RegisterHTTPv1Handler(grpSuppliers, "/{id}", http.MethodDelete)
		suppliers_get_list.RegisterHTTPv1Handler(grpSuppliers, "/", http.MethodGet)
		supplier_get.RegisterHTTPv1Handler(grpSuppliers, "/{id}", http.MethodGet)

		grpImages := huma.NewGroup(grpAuthNeed, "/images")
		image_update.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodPut)
		image_delete.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodDelete)
		image_get.RegisterHTTPv1Handler(grpImages, "/{id}", http.MethodGet)

		grpAuth := huma.NewGroup(grp_v1, "/auth")
		auth_login.RegisterHTTPv1Handler(grpAuth, "/auth", http.MethodPost)
		auth_register.RegisterHTTPv1Handler(grpAuth, "/register", http.MethodPost)
		auth_reset.RegisterHTTPv1Handler(grpAuth, "/reset", http.MethodPost)
	}
	return mux
}
