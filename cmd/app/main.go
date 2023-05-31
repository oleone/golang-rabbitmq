package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/oleone/golang-rabbitmq/internal/infra/drivers"
	"github.com/oleone/golang-rabbitmq/internal/infra/http_requests"
	"github.com/oleone/golang-rabbitmq/internal/infra/messagin"
	"github.com/oleone/golang-rabbitmq/internal/infra/repository"
	"github.com/oleone/golang-rabbitmq/internal/usecase"
)

func main() {
	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "ecommercex")
	defer mySqlDriver.Close()

	rabbitMqDriver := drivers.NewRabbitQMDriver("admin", "admin", "localhost", "5672")
	defer rabbitMqDriver.Close()

	rabbitMqChannel := messagin.NewRabbitMqChannel(rabbitMqDriver)
	defer rabbitMqChannel.Close()

	productRepository := repository.NewProductRepositoryMysql(mySqlDriver.DB)
	orderRepository := repository.NewOrderRepositoryMysql(mySqlDriver.DB)
	orderProductRepository := repository.NewOrderProductRepositoryMysql(mySqlDriver.DB)

	createProductUsecase := usecase.NewCreateProductUseCase(productRepository)
	listProductUsecase := usecase.NewListProductsUseCase(productRepository)
	createOrderUsecase := usecase.NewCreateOrderUseCase(orderRepository, productRepository, orderProductRepository)

	productHandlers := http_requests.NewProductHandlers(createProductUsecase, listProductUsecase)
	orderHandlers := http_requests.NewOrderHandlers(createOrderUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductHandler)
	r.Post("/orders", orderHandlers.CreateOrderHandler)

	http.ListenAndServe(":8000", r)
}
