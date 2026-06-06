package clients_get_list

import (
	"shopapi/internal/domain"
)

type InputGetAllClients struct {
	Limit  int64 `query:"limit,omitzero" doc:"number of clients to return"`
	Offset int64 `query:"offset,omitzero" doc:"number of clients to skip"`
}

type OutputGetAllClients struct {
	Body struct {
		Clients []*domain.Client `json:"clients" doc:"list of all clients"`
	}
}
