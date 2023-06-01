package usecase

import (
	"github.com/oleone/marketplacex/internal/entity"
)

type OrderProductInputDto struct {
	OrderID   string
	ProductID string
	Quantity  int
}

type OrderProductOutputDto struct {
	OrderID   string
	ProductID string
	Quantity  int
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
	orderProduct := entity.NewOrderProduct(input.OrderID, input.ProductID, input.Quantity)

	err := u.OrderProductRepository.Create(orderProduct)

	if err != nil {
		return nil, err
	}

	return &OrderProductOutputDto{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}, nil
}
