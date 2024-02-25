package repository

import (
	"context"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/internal/model"
	"github.com/ffajarpratama/go-chat/pkg/util"
	"github.com/google/uuid"
)

func (r *Repository) CreateChatroom(ctx context.Context, data *model.Chatroom) error {
	return r.db.WithContext(ctx).Create(&data).Error
}

func (r *Repository) FindAndCountChatroom(ctx context.Context, params *request.ListChatroomQuery) ([]*model.Chatroom, int64, error) {
	var res = make([]*model.Chatroom, 0)
	var cnt int64

	query := r.db.
		WithContext(ctx).
		Model(&model.Chatroom{}).
		Preload("Sender").
		Preload("Receiver").
		Preload("Post")

	if params.UserID != uuid.Nil {
		query = query.Where("(sender_id = ? OR receiver_id = ?)", params.UserID, params.UserID)
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

func (r *Repository) FindOneChatroom(ctx context.Context, query ...interface{}) (*model.Chatroom, error) {
	var res *model.Chatroom

	if err := r.db.
		WithContext(ctx).
		Model(&model.Chatroom{}).
		Preload("Sender").
		Preload("Receiver").
		Preload("Post").
		Where(query[0], query[1:]...).
		First(&res).
		Error; err != nil {
		return nil, err
	}

	return res, nil
}
