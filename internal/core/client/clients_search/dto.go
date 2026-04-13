package clients_search

import (
	"shopapi/internal/domain"
)

type InputClientsSearch struct {
	Body struct {
		Name    string `json:"name" doc:"name of the clients" maxLength:"100" example:"Ivan"`
		Surname string `json:"surname" doc:"surname of the clients" maxLength:"100" example:"Petrovich"`
	}
}

type OutputClientsSearch struct {
	Body struct {
		Clients []*domain.Client `json:"clients" doc:"founded clients"`
	}
}
