package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/oleone/golang-rabbitmq/internal/infra/http_requests"
	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/products")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := http_requests.NewProductHandlers(createProductUsecase, listProductUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductHandler)

	fmt.Println("Listening in port :8000")
	http.ListenAndServe(":8000", r)
}
