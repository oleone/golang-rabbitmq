package entity

import (
	"time"

	"github.com/google/uuid"
)

var OrderStatus = newOrderStatus()

type orderStatus struct {
	Approved       string
	Created        string
	Failed         string
	Pending        string
	PendingPayment string
}

func newOrderStatus() *orderStatus {
	return &orderStatus{
		Approved:       "approved",
		Created:        "created",
		Failed:         "failed",
		Pending:        "pending",
		PendingPayment: "pending_payment",
	}
}

type OrderRepository interface {
	Create(order *Order) error
	FindAll() ([]*Order, error)
}

type Order struct {
	ID         string
	Status     string
	OrderItems []OrderItem
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Amount     float64
}

type OrderItem struct {
	Quantity           int
	ProductName        string
	ProductID          string
	UnitCost           float64
	TotalCost          float64
	ShippingCost       float64
	DiscountPercentage float64
}

func NewOrder(orderItems []OrderItem) *Order {
	var amount float64

	for _, orderItem := range orderItems {
		amount += orderItem.TotalCost
	}

	return &Order{
		ID:         uuid.New().String(),
		Status:     OrderStatus.Created,
		OrderItems: orderItems,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Amount:     amount,
	}
}

func (o *Order) SetToPendingPayment() {
	o.Status = OrderStatus.PendingPayment
}
