package repository

import (
	"database/sql"

	"github.com/oleone/golang-rabbitmq/internal/entity"
)

type OrderRepositoryMysql struct {
	DB *sql.DB
}

func NewOrderRepositoryMysql(db *sql.DB) *OrderRepositoryMysql {
	return &OrderRepositoryMysql{DB: db}
}

func (r *OrderRepositoryMysql) Create(order *entity.Order) error {
	_, err := r.DB.Exec("insert into orders (amount, created_at, id, status, updated_at) values(?, ?, ?, ?, ?)", order.Amount, order.CreatedAt, order.ID, order.Status, order.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositoryMysql) FindAll() ([]*entity.Order, error) {
	rows, err := r.DB.Query("select amount, created_at, id, products, status, updated_at from orders")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order

	return orders, nil
}
