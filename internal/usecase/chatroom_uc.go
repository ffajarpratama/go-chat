package usecase

import (
	"context"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/internal/model"
	"github.com/google/uuid"
)

func (u *Usecase) CreateChatroom(ctx context.Context, req *request.CreateChatroomReq) error {
	data := &model.Chatroom{
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Topic:      req.Topic,
		PostID:     req.PostID,
	}

	return u.repo.CreateChatroom(ctx, data)
}

func (u *Usecase) FindAndCountChatroom(ctx context.Context, params *request.ListChatroomQuery) ([]*model.Chatroom, int64, error) {
	return u.repo.FindAndCountChatroom(ctx, params)
}

func (u *Usecase) FindOneChatroom(ctx context.Context, chatroomID uuid.UUID) (*model.Chatroom, error) {
	return u.repo.FindOneChatroom(ctx, "chatroom_id = ?", chatroomID)
}
