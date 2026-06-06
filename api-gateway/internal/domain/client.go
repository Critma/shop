package domain

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

type Client struct {
	ID               uuid.UUID  `json:"id" doc:"client ID"`
	Name             string     `json:"name" doc:"client name"`
	Surname          string     `json:"surname" doc:"client surname"`
	Birthday         time.Time  `json:"birthday" doc:"client birthday"`
	Gender           Gender     `json:"gender" validate:"max=1" doc:"client gender" enum:"M,F"`
	RegistrationDate time.Time  `json:"registrationDate" doc:"client registration date"`
	AddressID        *uuid.UUID `json:"addressID,omitempty" doc:"client address ID"`
	Address          *Address   `json:"address,omitempty" doc:"client address"`
}
