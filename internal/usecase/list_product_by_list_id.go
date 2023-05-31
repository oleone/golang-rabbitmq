package usecase

import (
	"fmt"

	"github.com/oleone/marketplacex/internal/entity"
)

type ListProductByListIdUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductByListIdUseCase(productRepository entity.ProductRepository) *ListProductByListIdUseCase {
	return &ListProductByListIdUseCase{ProductRepository: productRepository}
}

func (u *ListProductByListIdUseCase) Execute(listId []string) ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindByListId(listId)
	var productsOutput []*ListProductsOutputDto

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
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

	return productsOutput, nil
}
