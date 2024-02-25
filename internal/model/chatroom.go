package model

import (
	"github.com/ffajarpratama/go-chat/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Chatroom struct {
	ChatroomID uuid.UUID              `json:"chatroom_id" gorm:"primaryKey;default:gen_random_uuid()"`
	SenderID   uuid.UUID              `json:"sender_id"`
	ReceiverID uuid.UUID              `json:"receiver_id"`
	Topic      constant.ChatroomTopic `json:"topic"`
	PostID     *uuid.UUID             `json:"post_id"`
	CreatedAt  int                    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int                    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  soft_delete.DeletedAt  `json:"-" gorm:"column:deleted_at"`

	Sender   *User `json:"sender" gorm:"foreignKey:SenderID;references:UserID"`
	Receiver *User `json:"receiver" gorm:"foreignKey:ReceiverID;references:UserID"`
	Post     *Post `json:"post" gorm:"->"`
}

func (Chatroom) TableName() string {
	return "tr_chatroom"
}
