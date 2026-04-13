package domain

import "github.com/google/uuid"

type Supplier struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	AddressID   *uuid.UUID `json:"addressID,omitempty"`
	Address     *Address   `json:"address,omitempty"`
	PhoneNumber string     `json:"phoneNumber"`
}
