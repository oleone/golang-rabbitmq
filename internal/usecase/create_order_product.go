package usecase

import (
	"github.com/oleone/golang-rabbitmq/internal/entity"
)

type OrderProductInputDto struct {
	OrderID   string
	ProductID string
}

type OrderProductOutputDto struct {
	OrderID   string
	ProductID string
}

type CreateOrderProductUseCase struct {
	OrderProductRepository entity.OrderProductRepository
}

func NewOrderProductUseCase(repository entity.OrderProductRepository) *CreateOrderProductUseCase {
	return &CreateOrderProductUseCase{
		OrderProductRepository: repository,
	}
}

func (u *CreateOrderProductUseCase) Execute(input *OrderProductInputDto) (*OrderProductOutputDto, error) {
	orderProduct := entity.NewOrderProduct(input.OrderID, input.ProductID)

	err := u.OrderProductRepository.Create(orderProduct)

	if err != nil {
		return nil, err
	}

	return &OrderProductOutputDto{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
	}, nil
}
