package image_get

import (
	"github.com/google/uuid"
)

type InputImageGet struct {
	ImageID uuid.UUID `path:"id" json:"id" doc:"imageID to get" format:"uuid"`
}

type OutputImageGet struct {
	ContentType string `header:"Content-Type"`
	Body        []byte
}
