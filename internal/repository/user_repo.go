package repository

import (
	"context"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/internal/model"
)

func (r *Repository) FindAndCountUser(ctx context.Context, params *request.ListUserQuery) ([]*model.User, int64, error) {
	var res = make([]*model.User, 0)
	var cnt int64

	return res, cnt, nil
}
