package usecase_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oleone/marketplacex/internal/infra/drivers"
	"github.com/oleone/marketplacex/internal/infra/repository"
	"github.com/oleone/marketplacex/internal/usecase"
)

func TestCreateProductUseCase(t *testing.T) {
	t.Log("CreateProductUseCase test initialized")

	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "marketplacex")

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
