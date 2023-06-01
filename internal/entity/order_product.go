package entity

import "github.com/google/uuid"

type OrderProduct struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  int
}

type OrderProductRepository interface {
	Create(orderProduct *OrderProduct) error
	FindByOrderId(orderID string) ([]*OrderProduct, error)
	FindByProductId(productID string) ([]*OrderProduct, error)
}

func NewOrderProduct(orderID string, productID string, quantity int) *OrderProduct {
	return &OrderProduct{
		ID:        uuid.New().String(),
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
	}
}
