package model

import (
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type Post struct {
	PostID    uuid.UUID             `json:"post_id" gorm:"primaryKey;default:gen_random_uuid()"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	CreatedBy uuid.UUID             `json:"created_by"`
	CreatedAt int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Creator *User `json:"creator" gorm:"foreignKey:CreatedBy;references:UserID"`
}

func (Post) TableName() string {
	return "tr_post"
}
