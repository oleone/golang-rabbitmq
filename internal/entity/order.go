package entity

import (
	"time"

	"github.com/google/uuid"
)

var OrderStatus = newOrderStatus()

type orderStatus struct {
	Approved string
	Created  string
	Failed   string
	Pending  string
}

func newOrderStatus() *orderStatus {
	return &orderStatus{
		Approved: "approved",
		Created:  "created",
		Failed:   "failed",
		Pending:  "pending",
	}
}

type OrderRepository interface {
	Create(order *Order) error
	FindAll() ([]*Order, error)
}

type Order struct {
	ID        string
	Status    string
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
	Amount    float64
}

func NewOrder(products []Product) *Order {
	var amount float64

	for _, product := range products {
		amount += product.Price
	}

	return &Order{
		ID:        uuid.New().String(),
		Status:    OrderStatus.Created,
		Products:  products,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Amount:    amount,
	}
}
