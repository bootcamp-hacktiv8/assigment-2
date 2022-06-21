package entities

import "time"

type Order struct {
	ID           int       `json:"id" form:"id"`
	CustomerName string    `json:"customer_name" form:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at" form:"ordered_at"`
	Item         []Item    `json:"items" form:"items"`
}
