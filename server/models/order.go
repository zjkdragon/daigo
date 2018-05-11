package model

import (
	"time"
)

type OrderItem struct {
	ProductId string
	Price     int64
	Count     int32
}

type Order struct {
	Id         string
	TotalPrice int64
	TotalCount int32
	List       []OrderItem
	OrderTime  time.Time
}
