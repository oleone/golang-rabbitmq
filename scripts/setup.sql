# docker compose exec mysql bash
# mysql -uroot -p marketplacex
use markertplacex;

create table
    products(
        id VARCHAR(255),
        name VARCHAR(255),
        price FLOAT,
        category VARCHAR(255),
        sub_category VARCHAR(255),
        offer_percentage INT,
        quantity INT,
        reservad_quantity INT,
        primary key (id)
    );

create table
    orders(
        id VARCHAR(255),
        status VARCHAR(255),
        created_at DATE,
        updated_at DATE,
        amount FLOAT,
        primary key (id)
    );

create table
    order_products(
        id VARCHAR(255),
        order_id VARCHAR(255),
        product_id VARCHAR(255),
        quantity INT,
        primary key(id),
        foreign key (order_id) REFERENCES orders(id),
        foreign key (product_id) REFERENCES products(id)
    );