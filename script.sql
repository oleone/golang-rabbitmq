# docker compose exec mysql bash
# mysql -uroot -p products
create table
    products (
        id VARCHAR(255),
        name VARCHAR(255),
        price FLOAT
    );