package client_create

import (
	"context"
	"errors"
	"shopapi/internal/domain"
	"time"
)

type Store interface {
	ClientCreate(ctx context.Context, name, surname string, birthday time.Time, gender domain.Gender) (*domain.Client, error)
}

type useCase struct {
	store Store
}

func NewUsecase(store Store) *useCase {
	return &useCase{
		store: store,
	}
}

func (u *useCase) ClientCreate(ctx context.Context, input *InputClientCreate) (*OutputClientCreate, error) {
	createdClient, err := u.store.ClientCreate(ctx, input.Body.Name, input.Body.Surname, input.Body.Birthday, input.Body.Gender)
	if err != nil {
		return nil, errors.New("Internal server error")
	}
	output := &OutputClientCreate{}
	output.Body.ID = createdClient.ID
	return output, nil
}
