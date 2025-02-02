-- +goose Up
-- +goose StatementBegin
create schema shop;

create table if not exists shop.Users
(
    id       serial primary key,
    name     varchar(255),
    email    varchar(255) unique,
    password varchar(255)
);

create table if not exists shop.Orders
(
    id           serial primary key,
    user_id      int,
    order_date   timestamp,
    total_amount float,
    constraint fk_orders_users foreign key (user_id) references shop.Users (id)
);
create index idx_order_user_id ON shop.Orders USING btree (user_id);

create table if not exists shop.Products
(
    id    serial primary key,
    name  varchar(255),
    price float
);

create table if not exists shop.OrderProducts
(
    order_id   int,
    product_id int,
    primary key (order_id, product_id),
    constraint fk_orderProduct_order foreign key (order_id) references shop.Orders on delete cascade,
    constraint fk_orderProduct_product foreign key (product_id) references shop.Products on delete cascade

);
create index idx_orderproducts_order_id ON shop.OrderProducts USING btree (order_id);
create index idx_orderproducts_products_id ON shop.OrderProducts USING btree (product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists shop.OrderProducts;
drop table if exists shop.Products;
drop table if exists shop.Orders;
drop table if exists shop.Users;
drop schema if exists shop;
-- +goose StatementEnd
