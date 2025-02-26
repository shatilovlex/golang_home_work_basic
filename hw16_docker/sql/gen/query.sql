-- name: UserCreate :one
INSERT INTO pg_storage.shop.users (name, email, password)
VALUES ($1, $2, $3) returning id;

-- name: Users :many
select * from pg_storage.shop.users limit $1 offset $2;

-- name: UpdateUserName :one
update pg_storage.shop.users set name = $2 where id = $1 returning id;

-- name: ProductCreate :one
INSERT INTO pg_storage.shop.products (name, price)
VALUES ($1, $2) returning id;

-- name: Products :many
select * from pg_storage.shop.products limit $1 offset $2;

-- name: UpdateProduct :one
update pg_storage.shop.products set name = $2, price = $3 where id = $1 returning id;

-- name: CreateOrder :one
INSERT INTO pg_storage.shop.orders (user_id, order_date, total_amount)
VALUES ($1, $2, $3) returning id;

-- name: CreateOrderProduct :exec
INSERT INTO pg_storage.shop.orderProducts (order_id, product_id)
VALUES ($1, $2);