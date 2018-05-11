package db

import "time"

type Order struct {
	Id         string `gorm:"primary_key"`
	TotalPrice int64
	TotalCount int32
	CreatedAt  time.Time
}
