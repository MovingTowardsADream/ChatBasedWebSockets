package v1

import (
	"ChatBasedWebSockets/internal/entity"
	"context"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (string, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}
