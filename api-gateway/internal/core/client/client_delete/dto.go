package client_delete

import (
	"github.com/google/uuid"
)

type InputClientDelete struct {
	ClientID uuid.UUID `path:"id" json:"id" doc:"clientID to delete" format:"uuid"`
}
