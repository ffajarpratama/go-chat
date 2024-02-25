package request

import (
	"github.com/ffajarpratama/go-chat/pkg/constant"
	base_request "github.com/ffajarpratama/go-chat/pkg/http/request"
	"github.com/google/uuid"
)

type CreateChatroomReq struct {
	SenderID   uuid.UUID              `json:"sender_id" validate:"required"`
	ReceiverID uuid.UUID              `json:"receiver_id" validate:"required"`
	Topic      constant.ChatroomTopic `json:"topic" validate:"required,oneof=post direct"`
	PostID     *uuid.UUID             `json:"post_id"`
}

type ListChatroomQuery struct {
	base_request.Query
	Topic  constant.ChatroomTopic
	PostID uuid.UUID
	UserID uuid.UUID
}
