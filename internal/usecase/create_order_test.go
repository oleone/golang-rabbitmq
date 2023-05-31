package usecase_test

import (
	"testing"

	"github.com/oleone/golang-rabbitmq/internal/infra/drivers"
	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func TestCreateOrder(t *testing.T) {

	t.Log("TestCreateOrder initialized")

	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "ecommercex")

	createOrderRepository := repository.NewOrderRepositoryMysql(mySqlDriver.DB)
	productRepository := repository.NewProductRepositoryMysql(mySqlDriver.DB)
	orderProductRepository := repository.NewOrderProductRepositoryMysql(mySqlDriver.DB)

	createOrderUsecase := usecase.NewCreateOrderUseCase(createOrderRepository, productRepository, orderProductRepository)
	createProductUsecase := usecase.NewCreateProductUseCase(productRepository)

	var products []*usecase.ListProductsOutputDto

	productInput := usecase.CreateProductInputDto{
		Name:            "Product",
		Price:           25.80,
		Category:        "Caterogy",
		Subcategory:     "Subcategory",
		OfferPercentage: 0,
		Quantity:        265,
	}
	var productListId []string

	for i := 0; i <= 4; i++ {
		productCreatedOutput, err := createProductUsecase.Execute(productInput)

		if err != nil {
			t.Fail()
			t.Log(err)
		}

		productListId = append(productListId, productCreatedOutput.ID)

		products = append(products, &usecase.ListProductsOutputDto{
			ID:               productCreatedOutput.ID,
			Name:             productCreatedOutput.Name,
			Price:            productCreatedOutput.Price,
			Category:         productCreatedOutput.Category,
			Subcategory:      productCreatedOutput.Subcategory,
			OfferPercentage:  productCreatedOutput.OfferPercentage,
			Quantity:         productCreatedOutput.Quantity,
			ReservadQuantity: productCreatedOutput.ReservadQuantity,
		})
	}

	input := usecase.CreateOrderInputDto{
		Products: products,
	}

	_, err := createOrderUsecase.Execute(input)

	if err != nil {
		t.Fail()
		t.Log(err)
	} else {
		t.Log("TestCreateOrder finality with success")
	}

}
