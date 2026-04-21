package image_update

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type InputImageUpdate struct {
	ImageID uuid.UUID `path:"id" json:"id" doc:"imageID to update" format:"uuid"`
	RawBody huma.MultipartFormFiles[struct {
		Image huma.FormFile `form:"image" contentType:"application/octet-stream" doc:"image to assign"`
	}]
}

type OutputImageUpdate struct {
	Body struct {
		Success bool `json:"success" doc:"indicates if the update was successful"`
	}
}
