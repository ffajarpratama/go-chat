package model

import (
	"github.com/ffajarpratama/go-chat/pkg/constant"
	"github.com/google/uuid"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	UserID    uuid.UUID             `json:"user_id" gorm:"primaryKey; default:gen_random_uuid()"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Password  string                `json:"-"`
	Role      constant.UserRole     `json:"role"`
	CreatedAt int                   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int                   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "tr_user"
}
