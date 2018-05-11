package db

import "time"

type Product struct {
	Id          string `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(20);not null"`
	Price       int64
	Description string `gorm:"type:varchar(256);not null"`
	CreatedAt   time.Time
}
