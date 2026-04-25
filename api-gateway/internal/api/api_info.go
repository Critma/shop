package api

import "github.com/danielgtaylor/huma/v2"

func getAdditionalInfo(info *huma.Info) *huma.Info {
	info.Description = "ShopAPI provide api for the shop of home appliances. Avaible tags: Clients, Images, Products, Suppliers, Auth."
	// info.License = &huma.License{
	// 	Name:       "MIT License",
	// 	Identifier: "MIT",
	// }
	return info
}
