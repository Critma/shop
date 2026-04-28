package clients_search

import (
	"context"
	"shopapi/internal/domain"
	"shopapi/pkg/render"
)

type Store interface {
	ClientsSearch(ctx context.Context, name, surname string) ([]*domain.Client, error)
}

type Usecase struct {
	store Store
}

var usecase *Usecase

func NewUsecase(store Store) *Usecase {
	uc := &Usecase{
		store: store,
	}

	usecase = uc

	return uc
}

func (u *Usecase) ClientsSearch(ctx context.Context, input *InputClientsSearch) (*OutputClientsSearch, error) {
	clients, err := u.store.ClientsSearch(ctx, input.Body.Name, input.Body.Surname)
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}

	output := &OutputClientsSearch{}
	output.Body.Clients = clients
	return output, nil
}
