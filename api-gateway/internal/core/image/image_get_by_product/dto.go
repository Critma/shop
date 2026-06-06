package image_get_by_product

import (
	"github.com/google/uuid"
)

type InputImageGetByProduct struct {
	ProductID uuid.UUID `path:"id" json:"id" doc:"productID to get image for" format:"uuid"`
}

type OutputImageGetByProduct struct {
	ContentType string `header:"Content-Type"`
	Body        []byte
}
