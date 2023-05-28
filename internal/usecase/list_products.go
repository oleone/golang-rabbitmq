package usecase

import "github.com/oleone/golang-rabbitmq/internal/entity"

type ListProductsOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
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
		return nil, err
	}

	var productsOutput []*ListProductsOutputDto

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}
