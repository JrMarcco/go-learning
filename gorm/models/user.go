package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"size:100;not null;index:idx_user_search"`
	Email     string         `json:"email" gorm:"size:255;uniqueIndex;not null"`
	Phone     string         `json:"phone" gorm:"size:20;index"`
	Status    int            `json:"status" gorm:"not null;default:1;index;comment:状态 1:active 2:inactive"`
	Age       int            `json:"age" gorm:"index"`
	City      string         `json:"city" gorm:"size:50;index"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Profile *UserProfile `json:"profile,omitempty" gorm:"foreignKey:UserID"`
	Roles   []Role       `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

// UserProfile 用户详情
type UserProfile struct {
	ID       uint       `json:"id" gorm:"primarykey"`
	UserID   uint       `json:"user_id" gorm:"not null;index"`
	Avatar   string     `json:"avatar" gorm:"size:255"`
	Bio      string     `json:"bio" gorm:"type:text"`
	Gender   int        `json:"gender" gorm:"comment:性别 1:male 2:female"`
	Birthday *time.Time `json:"birthday"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Role 角色模型
type Role struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"size:50;not null;uniqueIndex"`
	Description string         `json:"description" gorm:"size:255"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

func (Role) TableName() string {
	return "roles"
}
