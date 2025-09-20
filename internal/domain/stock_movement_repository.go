package domain

type StockMovementRepository interface {
	Create(movement *StockMovement) error
	ListByProductID(productID int) ([]*StockMovement, error)
}
