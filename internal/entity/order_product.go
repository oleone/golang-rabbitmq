package entity

import "github.com/google/uuid"

type OrderProduct struct {
	ID        string
	OrderID   string
	ProductID string
}

type OrderProductRepository interface {
	Create(orderProduct *OrderProduct) error
	FindByOrderId(orderID string) ([]*OrderProduct, error)
	FindByProductId(productID string) ([]*OrderProduct, error)
}

func NewOrderProduct(orderID string, productID string) *OrderProduct {
	return &OrderProduct{
		ID:        uuid.New().String(),
		OrderID:   orderID,
		ProductID: productID,
	}
}
