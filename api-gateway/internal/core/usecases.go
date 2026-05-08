package core

import (
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
)

func CreateUsecases(app config.App) {
	auth_header_middleware.NewUsecase(app.GrpcClient)

	client_create.NewUsecase(app.Store)
	client_delete.NewUsecase(app.Store)
	clients_search.NewUsecase(app.Store)
	clients_get_list.NewUsecase(app.Store)
	client_update_address.NewUsecase(app.Store)

	product_create.NewUsecase(app.Store)
	product_get.NewUsecase(app.Store)
	products_get_list.NewUsecase(app.Store)
	product_delete.NewUsecase(app.Store)

	image_assign.NewUsecase(app.Store)
	image_update.NewUsecase(app.Store)
	image_delete.NewUsecase(app.Store)
	image_get.NewUsecase(app.Store)
	image_get_by_product.NewUsecase(app.Store)

	supplier_create.NewUsecase(app.Store)
	supplier_update_address.NewUsecase(app.Store)
	supplier_delete.NewUsecase(app.Store)
	suppliers_get_list.NewUsecase(app.Store)
	supplier_get.NewUsecase(app.Store)

	auth_login.NewUsecase(app.GrpcClient)
	auth_register.NewUsecase(app.GrpcClient)
	auth_reset.NewUsecase(app.GrpcClient)
}
