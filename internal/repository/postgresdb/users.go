package postgresdb

import (
	"ChatBasedWebSockets/internal/entity"
	"ChatBasedWebSockets/pkg/postgres"
	"context"
)

type UserRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg: pg}
}

func (ur *UserRepo) GetAllUsers(ctx context.Context, userCh chan<- entity.User) {
	sqlQuery, args, _ := ur.pg.Builder.
		Select("id, email, username, password, created_at").
		From(tableUsers).
		ToSql()

	rows, _ := ur.pg.Pool.Query(ctx, sqlQuery, args...)

	go func() {
		defer rows.Close()

		for rows.Next() {
			var user entity.User

			_ = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.CreatedAt)

			userCh <- user
		}
		close(userCh)
	}()
}
