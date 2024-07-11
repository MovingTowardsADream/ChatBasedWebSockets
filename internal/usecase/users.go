package usecase

import (
	"ChatBasedWebSockets/internal/entity"
	"context"
	"log/slog"
)

type UsersWork interface {
	GetAllUsers(ctx context.Context, userCh chan<- entity.User)
}

type UsersUseCase struct {
	l         *slog.Logger
	usersWork UsersWork
}

func NewUsersUseCase(l *slog.Logger, uw UsersWork) *UsersUseCase {
	return &UsersUseCase{l: l, usersWork: uw}
}

func (uuc *UsersUseCase) GetAllUsers(ctx context.Context, userCh chan<- entity.User) {
	uuc.usersWork.GetAllUsers(ctx, userCh)
}
