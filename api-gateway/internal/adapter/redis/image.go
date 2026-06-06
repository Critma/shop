package redis

import (
	"context"
	"encoding/base64"
	"shopapi/internal/domain"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func getImageKey(id string) string {
	const prefix = "image:"
	return prefix + id
}

func (r *Redis) GetImage(ctx context.Context, id uuid.UUID) (*domain.Image, error) {
	ctx, cancel := context.WithTimeout(ctx, QUERY_TIMEOUT)
	defer cancel()

	cachedData, err := r.client.Get(ctx, getImageKey(id.String())).Result()
	if err != nil {
		switch err {
		case redis.Nil:
			return nil, domain.ErrNotFound
		default:
			log.Error().Err(err).Msg("Unable to get image from Redis")
			return nil, err
		}
	}

	decoded, err := base64.StdEncoding.DecodeString(cachedData)
	if err != nil {
		log.Error().Err(err).Msg("Unable to decode cached image")
		return nil, err
	}
	log.Info().Str("imageID", id.String()).Msg("Image found in Redis")
	return &domain.Image{
		ID:    id,
		Image: decoded,
	}, nil
}

func (r *Redis) SetImage(ctx context.Context, image *domain.Image) error {
	if err := validateImage(image); err != nil {
		return err
	}

	encodedData := base64.StdEncoding.EncodeToString(image.Image)
	if err := r.client.Set(ctx, getImageKey(image.ID.String()), encodedData, r.cfg.CacheTTL).Err(); err != nil {
		log.Error().Err(err).Msg("Unable to set image in Redis")
		return err
	}
	log.Info().Str("imageID", image.ID.String()).Msg("Image cached in Redis")
	return nil
}

func (r *Redis) DeleteImage(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return domain.ErrInvalidArgs
	}

	if err := r.client.Del(ctx, getImageKey(id.String())).Err(); err != nil {
		log.Error().Err(err).Msg("Unable to delete image from Redis")
		return err
	}
	log.Info().Str("imageID", id.String()).Msg("Image deleted from Redis")
	return nil
}

func validateImage(image *domain.Image) error {
	if image == nil {
		log.Error().Msg("Image to cache is nil")
		return domain.ErrInvalidArgs
	}
	if image.ID == uuid.Nil || image.Image == nil {
		log.Error().Any("image", image).Msg("Image to cache is invalid")
		return domain.ErrInvalidArgs
	}
	return nil
}
