package domain

type StockLevelRepository interface {
	Create(stockLevel *StockLevel) error
	GetByProductID(productID int) (*StockLevel, error)
	UpdateQuantity(stockLevel *StockLevel) error
}
