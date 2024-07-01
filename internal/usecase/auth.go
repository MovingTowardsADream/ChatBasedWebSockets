package usecase

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/repository/repository_error"
	"ChatBasedWebSockets/pkg/hasher"
	"ChatBasedWebSockets/pkg/logger"
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"time"
)

const (
	_defaultTimeout = 5 * time.Second
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	GetUser(ctx context.Context, username, password string) (entity.User, error)
}

type AuthUseCase struct {
	l    *slog.Logger
	auth Authorization

	signKey        string
	passwordHasher hasher.PasswordHasher
	tokenTTL       time.Duration
}

func NewAuthUseCase(log *slog.Logger, auth Authorization,
	signKey string, passwordHasher hasher.PasswordHasher, tokenTTL time.Duration,
) *AuthUseCase {
	return &AuthUseCase{
		l:              log,
		auth:           auth,
		signKey:        signKey,
		passwordHasher: passwordHasher,
		tokenTTL:       tokenTTL,
	}
}

func (auc *AuthUseCase) CreateUser(ctx context.Context, user entity.User) (string, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, _defaultTimeout)
	defer cancel()

	user.Password = auc.passwordHasher.Hash(user.Password)

	id, err := auc.auth.CreateUser(ctxTimeout, user)
	if err != nil {
		if errors.Is(err, ErrTimeout) {
			return "", ErrTimeout
		} else if errors.Is(err, repository_error.ErrAlreadyExists) {
			return "", ErrAlreadyExists
		}

		auc.l.Error("AuthUseCase.CreateUser - auc.auth.CreateUser: ", logger.Err(err))

		return "", err
	}

	return id, err
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (auc *AuthUseCase) GenerateToken(ctx context.Context, username, password string) (string, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, _defaultTimeout)
	defer cancel()

	user, err := auc.auth.GetUser(ctxTimeout, username, auc.passwordHasher.Hash(password))

	if err != nil {
		if errors.Is(err, ErrTimeout) {
			return "", ErrTimeout
		} else if errors.Is(err, repository_error.ErrNotFound) {
			return "", ErrNotFound
		}

		auc.l.Error("AuthUseCase.GenerateToken - auc.auth.GetUser: ", logger.Err(err))

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(auc.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(auc.signKey))
}

func (auc *AuthUseCase) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(auc.signKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
