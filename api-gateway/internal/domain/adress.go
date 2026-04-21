package domain

import "github.com/google/uuid"

type Address struct {
	ID      uuid.UUID `json:"id" doc:"address ID"`
	Country string    `json:"country" doc:"Full name country" maxLength:"100"`
	City    string    `json:"city" doc:"Full name city" maxLength:"100"`
	Street  string    `json:"street" doc:"Full name street" maxLength:"255"`
}
