package domain

import "time"

type StockMovement struct {
	ID           int
	ProductID    int
	MovementType string
	Quantity     int
	Reason       string
	UserID       string
	CreatedAt    time.Time
}
