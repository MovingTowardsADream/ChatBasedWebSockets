package postgresdb

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/internal/repository/repository_error"
	"ChatBasedWebSockets/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	tableUsers = "users"
)

type AuthRepo struct {
	pg *postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg: pg}
}

func (ar *AuthRepo) CreateUser(ctx context.Context, user entity.User) (string, error) {
	sql, args, _ := ar.pg.Builder.
		Insert(tableUsers).
		Columns("email", "username", "password").
		Values(user.Email, user.Username, user.Password).
		Suffix("RETURNING id").
		ToSql()

	var id string
	err := ar.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return "", repository_error.ErrAlreadyExists
			}
		}
		return "", fmt.Errorf("AuthRepo.CreateUser - r.Pool.QueryRow: %v", err)
	}

	return id, nil
}

func (ar *AuthRepo) GetUser(ctx context.Context, username, password string) (entity.User, error) {
	fmt.Println(username, password)
	sql, args, _ := ar.pg.Builder.
		Select("id, email, username, password, created_at").
		From(tableUsers).
		Where("username = ? AND password = ?", username, password).
		ToSql()

	var user entity.User
	err := ar.pg.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repository_error.ErrNotFound
		}
		return entity.User{}, fmt.Errorf("AuthRepo.GetUser - r.Pool.QueryRow: %v", err)
	}

	return user, nil
}
