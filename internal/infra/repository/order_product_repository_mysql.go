package repository

import (
	"database/sql"

	"github.com/oleone/marketplacex/internal/entity"
)

type OrderProductRepositoryMysql struct {
	DB *sql.DB
}

func NewOrderProductRepositoryMysql(db *sql.DB) *OrderProductRepositoryMysql {
	return &OrderProductRepositoryMysql{
		DB: db,
	}
}

func (r *OrderProductRepositoryMysql) Create(orderProduct *entity.OrderProduct) error {
	_, err := r.DB.Exec("insert into order_products (id, order_id, product_id) values (?, ?, ?)", orderProduct.ID, orderProduct.OrderID, orderProduct.ProductID)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderProductRepositoryMysql) FindByOrderId(orderID string) ([]*entity.OrderProduct, error) {
	rows, err := r.DB.Query("select id, order_id, product_id where order_id = ?", orderID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order_products []*entity.OrderProduct

	for rows.Next() {
		var order_product entity.OrderProduct
		err = rows.Scan(
			&order_product.ID,
			&order_product.OrderID,
			&order_product.ProductID,
		)
		order_products = append(order_products, &order_product)
	}

	if err != nil {
		return nil, err
	}

	return order_products, nil
}

func (r *OrderProductRepositoryMysql) FindByProductId(productID string) ([]*entity.OrderProduct, error) {
	rows, err := r.DB.Query("select id, order_id, product_id where order_id = ?", productID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order_products []*entity.OrderProduct

	for rows.Next() {
		var order_product entity.OrderProduct
		err = rows.Scan(
			&order_product.ID,
			&order_product.OrderID,
			&order_product.ProductID,
		)
		order_products = append(order_products, &order_product)
	}

	if err != nil {
		return nil, err
	}

	return order_products, nil
}
