package usecase

import (
	"context"

	"github.com/critma/auth/internal/config"
	"github.com/critma/auth/internal/domain"
	"github.com/critma/auth/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const BCRYPT_COST = bcrypt.DefaultCost

type AuthUsecase struct {
	Config *config.Config
	log    *zap.Logger
	Store  domain.Storage
}

func NewAuthUsecase(config *config.Config, log *zap.Logger, store domain.Storage) *AuthUsecase {
	uc := &AuthUsecase{
		Config: config,
		log:    log,
		Store:  store,
	}

	return uc
}

func (uc *AuthUsecase) TokenValidate(tokenString string) (bool, error) {
	ok, _, err := jwt.Parse(tokenString, uc.Config.Token.Secret)
	if err != nil {
		uc.log.Error("error parsing token", zap.Error(err))
		return false, err
	}

	return ok, nil
}

func (uc *AuthUsecase) UserRegister(ctx context.Context, user *domain.User) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password.Value), BCRYPT_COST)
	if err != nil {
		uc.log.Error("failed to generate password hash", zap.Error(err))
		return "", err
	}
	user.Password.Hash = passHash
	user, err = uc.Store.CreateUser(ctx, user)
	if err != nil {
		uc.log.Error("failed to save user", zap.Error(err))
		return "", err
	}

	tokenString, err := jwt.Sign(jwt.UserClaims{ID: user.ID, Email: user.Email}, uc.Config.Token.TTL, uc.Config.Token.Secret)
	if err != nil {
		uc.log.Error("failed to generate token", zap.Error(err))
		return "", err
	}
	return tokenString, nil
}

func (uc *AuthUsecase) UserLogin(ctx context.Context, email, password string) (string, error) {
	user, err := uc.Store.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(user.Password.Hash, []byte(password))
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	tokenString, err := jwt.Sign(jwt.UserClaims{ID: user.ID, Email: user.Email}, uc.Config.Token.TTL, uc.Config.Token.Secret)
	if err != nil {
		uc.log.Error("failed to generate token", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (uc *AuthUsecase) PasswordRecovery(ctx context.Context, email string) error {
	user, err := uc.Store.GetUser(ctx, email)
	if err != nil {
		return err
	}

	uc.log.Info("RECOVERY PASSWORD", zap.String("password_hash", string(user.Password.Hash)))
	return nil
}

func (uc *AuthUsecase) PasswordChange(ctx context.Context, email, oldPassword, newPassword string) error {
	user, err := uc.Store.GetUser(ctx, email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password.Hash, []byte(oldPassword))
	if err != nil {
		return domain.ErrInvalidCredentials
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), BCRYPT_COST)
	if err != nil {
		uc.log.Error("failed to generate password hash", zap.Error(err))
		return err
	}
	if err := uc.Store.UpdateUserPasswordHash(ctx, user.ID, string(passHash)); err != nil {
		return err
	}
	return nil
}
