package domain

type StockReservationRepository interface {
	Create(reservation *StockReservation) error
	GetByID(id int) (*StockReservation, error)
	ListActiveByProduct(productID int) ([]*StockReservation, error)
	UpdateStatus(id int, status string) error
}
