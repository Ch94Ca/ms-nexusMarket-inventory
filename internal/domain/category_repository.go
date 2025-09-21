package domain

type CategoryRepository interface {
	Create(category *Category) error
	GetByID(id int) (*Category, error)
	ListAll() ([]*Category, error)
	Update(category *Category) error
	Delete(id int) error
}
