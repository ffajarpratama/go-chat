package usecase

import (
	"context"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/internal/model"
)

func (u *Usecase) FindAndCountUser(ctx context.Context, params *request.ListUserQuery) ([]*model.User, int64, error) {
	return u.repo.FindAndCountUser(ctx, params)
}
