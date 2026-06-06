package domain

import "github.com/google/uuid"

type ImageData []byte

type Image struct {
	ID    uuid.UUID `json:"id"`
	Image ImageData `json:"image"`
}
