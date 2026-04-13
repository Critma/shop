package image_delete

import (
	"github.com/google/uuid"
)

type InputImageDelete struct {
	ImageID uuid.UUID `path:"id" json:"id" doc:"imageID to delete" format:"uuid"`
}