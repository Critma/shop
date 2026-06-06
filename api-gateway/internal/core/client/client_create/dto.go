package client_create

import (
	"shopapi/internal/domain"
	"time"

	"github.com/google/uuid"
)

type InputClientCreate struct {
	Body struct {
		Name     string        `json:"name" minLength:"2" maxLength:"100" doc:"Client name" example:"Ivan"`
		Surname  string        `json:"surname" minLength:"5" maxLength:"100" doc:"Client surname" example:"Petrovich"`
		Birthday time.Time     `json:"birthday" doc:"Client birthday"`
		Gender   domain.Gender `json:"gender" maxLength:"1" doc:"Client Gender" example:"M" enum:"M,F"`
	}
}

type OutputClientCreate struct {
	Body struct {
		ID uuid.UUID `json:"id" format:"uuid" doc:"client ID"`
	}
}
