package usecase

import (
	"fmt"

	"github.com/oleone/golang-rabbitmq/internal/entity"
)

type ListProductsOutputDto struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
	Subcategory     string  `json:"sub_category"`
	OfferPercentage float64 `json:"offer_percentage"`
	Quantity        int     `json:"quantity"`
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var productsOutput []*ListProductsOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:              product.ID,
			Name:            product.Name,
			Price:           product.Price,
			Category:        product.Category,
			Subcategory:     product.Subcategory,
			OfferPercentage: product.OfferPercentage,
			Quantity:        product.Quantity,
		})
	}

	return productsOutput, nil
}
