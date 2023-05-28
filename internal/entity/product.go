package entity

import "github.com/google/uuid"

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

type Product struct {
	ID               string
	Name             string
	Price            float64
	Category         string
	Subcategory      string
	OfferPercentage  *float64
	Quantity         int
	ReservadQuantity int
}

func NewProduct(
	name string,
	price float64,
	category string,
	subcategory string,
	offerPercentage *float64,
	quantity int,
) *Product {
	return &Product{
		ID:              uuid.New().String(),
		Name:            name,
		Price:           price,
		Category:        category,
		Subcategory:     subcategory,
		OfferPercentage: offerPercentage,
		Quantity:        quantity,
	}
}
