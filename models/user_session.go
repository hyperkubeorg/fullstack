package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	Token     string `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID string `gorm:"type:text;not null"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}

func (b *UserSession) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (b *UserSession) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Token == "" {
		b.Token, _ = randomHex(128)
	}
	return nil
}

func (b *UserSession) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
