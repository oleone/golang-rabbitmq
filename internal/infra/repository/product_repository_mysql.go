package repository

import (
	"database/sql"

	"github.com/oleone/golang-rabbitmq/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("insert into products (id, name, price, category, sub_category, offer_percentage, quantity, reservad_quantity) values (?, ?, ?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Price, product.Category, product.Subcategory, product.OfferPercentage, product.Quantity, product.ReservadQuantity)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price, category, sub_category, offer_percentage, quantity, reservad_quantity from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Category,
			&product.Subcategory,
			&product.OfferPercentage,
			&product.Quantity,
			&product.ReservadQuantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (r *ProductRepositoryMysql) FindByListId(listId []string) ([]*entity.Product, error) {

	var list string
	for i, id := range listId {
		list += "'" + id + "'"

		if i < len(listId)-1 {
			list += ","
		}

	}

	rows, err := r.DB.Query("select id, name, price, category, sub_category, offer_percentage, quantity, reservad_quantity from products where id in (" + list + ")")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Category,
			&product.Subcategory,
			&product.OfferPercentage,
			&product.Quantity,
			&product.ReservadQuantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
