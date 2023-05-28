# docker compose exec mysql bash
# mysql -uroot -p products
create table
    products (
        id VARCHAR(255),
        name VARCHAR(255),
        price FLOAT,
        category VARCHAR(255),
        sub_category VARCHAR(255),
        offer_percentage INT,
        quantity INT,
        reservad_quantity INT
    );