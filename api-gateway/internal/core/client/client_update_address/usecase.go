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

func (u *Usecase) UpdateClientAddress(ctx context.Context, input *InputUpdateClientAddress) (*OutputUpdateClientAddress, error) {
	err := u.store.ClientChangeAddress(ctx, input.ID, toAddress(input))
	if err != nil {
		return nil, render.ErrLeadInternalServer
	}

	output := &OutputUpdateClientAddress{}
	output.Body.Success = true
	return output, nil
}
