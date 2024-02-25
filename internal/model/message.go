package model

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Message struct {
	MessageID  uuid.UUID             `json:"message_id" gorm:"primaryKey; default:gen_random_uuid()"`
	ChatroomID uuid.UUID             `json:"chatroom_id"`
	UserID     uuid.UUID             `json:"user_id"`
	ReadAt     int                   `json:"read_at"`
	CreatedAt  int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Chatroom *Chatroom `json:"chatroom" gorm:"->"`
	User     *User     `json:"user" gorm:"->"`
}

func (Message) TableName() string {
	return "tr_message"
}
