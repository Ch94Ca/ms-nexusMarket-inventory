package domain

import "time"

type StockLevel struct {
	ProductID int
	Quantity  int
	UpdatedAt time.Time
}
