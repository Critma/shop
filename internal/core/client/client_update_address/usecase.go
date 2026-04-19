package client_update_address

import (
	"context"
	"shopapi/internal/domain"
	"shopapi/pkg/render"

	"github.com/google/uuid"
)

type Store interface {
	ClientChangeAddress(ctx context.Context, clientID uuid.UUID, address *domain.Address) error
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) UpdateClientAddress(ctx context.Context, input *InputUpdateClientAddress) (*OutputUpdateClientAddress, error) {
	err := u.store.ClientChangeAddress(ctx, input.ID, toAddress(input))
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}

	output := &OutputUpdateClientAddress{}
	output.Body.Success = true
	return output, nil
}
