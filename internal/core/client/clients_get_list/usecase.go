package clients_get_list

import (
	"context"
	"shopapi/internal/adapter/postgres"
	"shopapi/internal/domain"
	"shopapi/pkg/render"
)

type Store interface {
	ClientGetList(ctx context.Context, params postgres.ListParams) ([]*domain.Client, error)
}

type useCase struct {
	store Store
}

func newUseCase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) GetAllClients(ctx context.Context, input *InputGetAllClients) (*OutputGetAllClients, error) {
	var params postgres.ListParams
	// skip zero values
	if input.Limit != 0 {
		params.Limit = &input.Limit
	}
	if input.Offset != 0 {
		params.Offset = &input.Offset
	}

	clients, err := u.store.ClientGetList(ctx, params)
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}

	output := &OutputGetAllClients{}
	output.Body.Clients = clients
	return output, nil
}
