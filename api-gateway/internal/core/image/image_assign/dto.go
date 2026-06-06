package image_assign

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type InputImageAssign struct {
	ProductID uuid.UUID `path:"id" format:"uuid" doc:"product ID"`
	RawBody   huma.MultipartFormFiles[struct {
		Image huma.FormFile `form:"image" contentType:"application/octet-stream" doc:"image to assign"`
	}]
}

type OutputImageAssign struct {
	Body struct {
		ImageID uuid.UUID `json:"imageID" doc:"created Image ID"`
	}
}
