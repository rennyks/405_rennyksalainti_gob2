package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique;type:varchar(191)"`
	Brand     string `gorm:"not null;unique;type:varchar(191)"`
	UserID    uint   // Foreign Key ID from User
	CreatedAt time.Time
	UpdateAt  time.Time
}	

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("product before create()")

	if len(p.Name) < 4 {
		err = errors.New("product name is too short")
	}

	return
}