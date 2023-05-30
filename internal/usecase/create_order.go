package usecase

import (
	"fmt"
	"time"

	"github.com/oleone/golang-rabbitmq/internal/entity"
)

type CreateOrderInputDto struct {
	Products []*ListProductsOutputDto
}

type CreateOrderOutputDto struct {
	ID        string
	Status    string
	Products  []entity.Product
	CreatedAt time.Time
	UpdatedAt time.Time
	Amount    float64
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

func (u *CreateOrderUseCase) Execute(input CreateOrderInputDto, productListId []string) (*CreateOrderOutputDto, error) {
	listProductByIdUsecase := NewListProductByListIdUseCase(u.ProductRepository)
	orderProductUsecase := NewOrderProductUseCase(u.OrderProductRepository)

	listProductsOutput, err := listProductByIdUsecase.Execute(productListId)
	var listProducts []entity.Product

	for _, product := range listProductsOutput {
		fmt.Println(product.ID)
		listProducts = append(listProducts, entity.Product{
			ID:               product.ID,
			Name:             product.Name,
			Price:            product.Price,
			Category:         product.Category,
			Subcategory:      product.Subcategory,
			OfferPercentage:  product.OfferPercentage,
			Quantity:         product.Quantity,
			ReservadQuantity: product.ReservadQuantity,
		})
	}
	order := entity.NewOrder(listProducts)
	err = u.OrderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	for _, product := range listProducts {
		var orderProduct = OrderProductInputDto{
			OrderID:   order.ID,
			ProductID: product.ID,
		}
		orderProductUsecase.Execute(&orderProduct)
	}

	return &CreateOrderOutputDto{
		ID:        order.ID,
		Status:    order.Status,
		Products:  order.Products,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Amount:    order.Amount,
	}, nil
}
