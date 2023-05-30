package entity_test

import (
	"testing"

	"github.com/oleone/golang-rabbitmq/internal/entity"
)

func TestNewProduct(t *testing.T) {
	t.Log("TestNewProduct test inicializaded")

	product := entity.NewProduct("Product 1", 152.85, "Category 1", "Subcategory 1", 0, 52)

	if product == nil {
		t.Fail()
	}

	t.Log("TestNewProduct test successed")
}
