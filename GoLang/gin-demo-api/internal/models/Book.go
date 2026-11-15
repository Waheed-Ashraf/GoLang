package models

import (
	"time"
	"gorm.io/gorm"
)

type Book struct { 
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"size:200; not null"`
	Author string `gorm:"size:100;not null;index"`
	Price float64 `gorm:"numeric(10 , 2); not null"`
	ISBN string `gorm:"size:20; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}