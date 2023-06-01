package usecase

import (
	"errors"
	"time"

	"github.com/oleone/marketplacex/internal/entity"
)

type CreateOrderInputDto struct {
	OrderItems []*OrderItemInputDto
}

type CreateOrderOutputDto struct {
	ID         string               `json:"id"`
	Status     string               `json:"status"`
	OrderItems []OrderItemOutputDto `json:"order_items"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
	Amount     float64              `json:"amount"`
}

type OrderItemInputDto struct {
	ProductID    string  `json:"product_id"`
	Quantity     int     `json:"quantity"`
	ShippingCost float64 `json:"shipping_cost"`
}

type OrderItemOutputDto struct {
	ProductID    string  `json:"product_id"`
	Quantity     int     `json:"quantity"`
	ShippingCost float64 `json:"shipping_cost"`
}

type CreateOrderUseCase struct {
	OrderRepository        entity.OrderRepository
	ProductRepository      entity.ProductRepository
	OrderProductRepository entity.OrderProductRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository, productRepository entity.ProductRepository, orderProductRepository entity.OrderProductRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:        orderRepository,
		ProductRepository:      productRepository,
		OrderProductRepository: orderProductRepository,
	}
}

func (u *CreateOrderUseCase) Execute(input CreateOrderInputDto) (*CreateOrderOutputDto, error) {
	listProductByIdUsecase := NewListProductByListIdUseCase(u.ProductRepository)
	orderProductUsecase := NewOrderProductUseCase(u.OrderProductRepository)

	var productListId []string

	for _, orderItems := range input.OrderItems {
		productListId = append(productListId, orderItems.ProductID)
	}

	listProductsOutput, err := listProductByIdUsecase.Execute(productListId)

	if len(listProductsOutput) == 0 {
		return nil, errors.New("ProductIDs is not found")
	}

	var listOrderItems []entity.OrderItem

	for _, product := range listProductsOutput {
		listOrderItems = append(listOrderItems, entity.OrderItem{
			ProductID:          product.ID,
			ProductName:        product.Name,
			Quantity:           2,
			UnitCost:           25.5,
			TotalCost:          28,
			ShippingCost:       15,
			DiscountPercentage: 3,
		})
	}
	order := entity.NewOrder(listOrderItems)
	err = u.OrderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	for _, orderItem := range listOrderItems {
		var orderProduct = OrderProductInputDto{
			OrderID:   order.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
		}
		orderProductUsecase.Execute(&orderProduct)
	}

	var orderItemsOutput []OrderItemOutputDto

	for _, lOrderItem := range listOrderItems {
		orderItemsOutput = append(orderItemsOutput, OrderItemOutputDto{
			ProductID:    lOrderItem.ProductID,
			Quantity:     lOrderItem.Quantity,
			ShippingCost: lOrderItem.ShippingCost,
		})
	}

	return &CreateOrderOutputDto{
		ID:         order.ID,
		Status:     order.Status,
		OrderItems: orderItemsOutput,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Amount:     order.Amount,
	}, nil
}
