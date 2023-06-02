package usecase

import (
	"github.com/oleone/marketplacex/internal/entity"
)

type CreateProductInputDto struct {
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
	Subcategory     string  `json:"sub_category"`
	OfferPercentage float64 `json:"offer_percentage"`
	Quantity        int     `json:"quantity"`
}

type CreateProductOutputDto struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Category         string  `json:"category"`
	Subcategory      string  `json:"sub_category"`
	OfferPercentage  float64 `json:"offer_percentage"`
	Quantity         int     `json:"quantity"`
	ReservadQuantity int     `json:"reservad_quantity"`
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(
		input.Name,
		input.Price,
		input.Category,
		input.Subcategory,
		input.OfferPercentage,
		input.Quantity,
	)

	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDto{
		ID:               product.ID,
		Name:             product.Name,
		Price:            product.Price,
		Category:         product.Category,
		Subcategory:      product.Subcategory,
		OfferPercentage:  product.OfferPercentage,
		Quantity:         product.Quantity,
		ReservadQuantity: product.ReservadQuantity,
	}, nil
}
