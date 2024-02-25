package repository

import (
	"context"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/internal/model"
	"github.com/ffajarpratama/go-chat/pkg/util"
)

func (r *Repository) FindAndCountUser(ctx context.Context, params *request.ListUserQuery) ([]*model.User, int64, error) {
	var res = make([]*model.User, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.User{})

	if params.Keyword != "" {
		query = query.Where("name ILIKE ?", "%"+params.Keyword+"%")
	}

	if params.Role != "" {
		query = query.Where("role = ?", params.Role)
	}

	if err := query.Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	if params.Sort != "" {
		query = query.Order(util.TransformSortClause("created_at", params.Sort))
	}

	if err := query.
		Limit(params.Limit).
		Offset(util.CalculateOffset(params.Page, params.Limit)).
		Find(&res).
		Error; err != nil {
		return nil, 0, err
	}

	return res, cnt, nil
}
