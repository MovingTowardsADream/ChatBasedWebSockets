package postgresdb

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/pkg/postgres"
	"context"
)

type AuthRepo struct {
	pg *postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg: pg}
}

func (ar *AuthRepo) CreateUser(ctx context.Context, user entity.User) (string, error) {
	return "", nil
}

func (ar *AuthRepo) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	return entity.User{}, nil
}
