package usecase_test

import (
	"testing"

	"github.com/oleone/marketplacex/internal/infra/drivers"
	"github.com/oleone/marketplacex/internal/infra/repository"
	"github.com/oleone/marketplacex/internal/usecase"
)

func TestCreateOrder(t *testing.T) {

	t.Log("TestCreateOrder with ProductIDs not found")

	mySqlDriver := drivers.NewMySqlDriver("root", "root", "localhost", "3306", "marketplacex")
	defer mySqlDriver.Close()

	createOrderRepository := repository.NewOrderRepositoryMysql(mySqlDriver.DB)
	productRepository := repository.NewProductRepositoryMysql(mySqlDriver.DB)
	orderProductRepository := repository.NewOrderProductRepositoryMysql(mySqlDriver.DB)

	createOrderUsecase := usecase.NewCreateOrderUseCase(createOrderRepository, productRepository, orderProductRepository)

	productListId := []string{"1585222"}

	var orderItems []*usecase.OrderItemInputDto

	for i := 0; i <= len(productListId)-1; i++ {
		orderItems = append(orderItems, &usecase.OrderItemInputDto{
			ProductID:    productListId[i],
			Quantity:     4,
			ShippingCost: 25.8,
		})
	}

	input := usecase.CreateOrderInputDto{
		OrderItems: orderItems,
	}

	_, err := createOrderUsecase.Execute(input)

	if err != nil {
		t.Log("TestCreateOrder with ProductIDs not found finality with success")
	} else {
		t.Fail()
		t.Log(err)
	}

}
