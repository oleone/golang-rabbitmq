package usecase_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func TestCreateProductUseCase(t *testing.T) {
	t.Log("CreateProductUseCase test initialized")

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommercex")

	if err != nil {
		t.Fail()
		t.Log(err)
	}
	defer db.Close()

	input := usecase.CreateProductInputDto{
		Name:            "Product 1",
		Category:        "Category 1",
		Subcategory:     "Subcategory 1",
		OfferPercentage: 0,
		Quantity:        254,
		Price:           152.85,
	}

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)

	_, err = createProductUsecase.Execute(input)

	if err != nil {
		t.Fail()
		t.Log(err)
	}

	t.Log("CreateProductUseCase test finalizad with success")
}
