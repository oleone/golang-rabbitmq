package usecase_test

import (
	"database/sql"
	"testing"

	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func TestCreateOrder(t *testing.T) {

	t.Log("TestCreateOrder initialized")

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommercex")

	if err != nil {
		t.Fail()
		t.Log(err)
	}
	defer db.Close()

	createOrderRepository := repository.NewOrderRepositoryMysql(db)
	productRepository := repository.NewProductRepositoryMysql(db)
	orderProductRepository := repository.NewOrderProductRepositoryMysql(db)

	orderUsecase := usecase.NewCreateOrderUseCase(createOrderRepository, productRepository, orderProductRepository)
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
	var productCreatedOutput *usecase.CreateProductOutputDto
	var productListId []string

	for i := 0; i <= 4; i++ {
		productCreatedOutput, err = createProductUsecase.Execute(productInput)

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

	_, err = orderUsecase.Execute(input, productListId)

	if err != nil {
		t.Fail()
		t.Log(err)
	} else {
		t.Log("TestCreateOrder finality with success")
	}

}
