package usecase_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oleone/golang-rabbitmq/internal/infra/drivers"
	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func TestCreateProductUseCase(t *testing.T) {
	t.Log("CreateProductUseCase test initialized")

	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "ecommercex")

	input := usecase.CreateProductInputDto{
		Name:            "Product 1",
		Category:        "Category 1",
		Subcategory:     "Subcategory 1",
		OfferPercentage: 0,
		Quantity:        254,
		Price:           152.85,
	}

	repository := repository.NewProductRepositoryMysql(mySqlDriver.DB)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)

	_, err := createProductUsecase.Execute(input)

	if err != nil {
		t.Fail()
		t.Log(err)
	}

	t.Log("CreateProductUseCase test finalizad with success")
}
