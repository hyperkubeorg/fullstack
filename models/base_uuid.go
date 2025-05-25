package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseUUID struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseUUID) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (b *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

func (b *BaseUUID) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}

type EphemeralBaseUUID struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *EphemeralBaseUUID) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (b *EphemeralBaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

func (b *EphemeralBaseUUID) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
