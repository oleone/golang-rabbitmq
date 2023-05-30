package http_requests

import "github.com/oleone/golang-rabbitmq/internal/usecase"

type OrderHandlers struct {
	OrderUseCase usecase.CreateOrderUseCase
}

type NewOrderHandlers struct {
}
