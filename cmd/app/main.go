package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/oleone/golang-rabbitmq/internal/infra/drivers"
	"github.com/oleone/golang-rabbitmq/internal/infra/http_requests"
	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func main() {
	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "ecommercex")

	repository := repository.NewProductRepositoryMysql(mySqlDriver.DB)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := http_requests.NewProductHandlers(createProductUsecase, listProductUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductHandler)

	http.ListenAndServe(":8000", r)
}
