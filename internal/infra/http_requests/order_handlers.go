package http_requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oleone/marketplacex/internal/usecase"
)

type OrderHandlers struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderHandlers(createOrderUsecase *usecase.CreateOrderUseCase) *OrderHandlers {
	return &OrderHandlers{
		CreateOrderUseCase: createOrderUsecase,
	}
}

func (h *OrderHandlers) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	var input usecase.CreateOrderInputDto

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateOrderUseCase.Execute(input)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
