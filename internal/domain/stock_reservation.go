package domain

import "time"

type StockReservation struct {
	ID          int
	ProductID   int
	ReservedQty int
	ReferenceID string
	ReservedAt  time.Time
	ExpiresAt   *time.Time
	Status      string
	UserID      string
}
