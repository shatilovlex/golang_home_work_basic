drop schema if exists shop cascade;
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

INSERT INTO shop.Users (name, email, password)
VALUES ('Вася', 'vasia@mail.loc', 'hash'),
       ('Петр', 'petya@mail.loc', 'hash'),
       ('Роман', 'roma@mail.loc', 'hash');

INSERT INTO shop.Products (id, name, price)
VALUES (1, 'Монополия', 10),
       (2, 'НЭП', 40),
       (3, 'Шахматы', 80),
       (4, 'Крокодил', 5.5)
;

INSERT INTO shop.Orders (id, user_id, order_date, total_amount)
VALUES (1,
        1,
        to_timestamp('2025-01-01 00:00:00', 'yyyy-mm-dd hh24:mi:ss'),
        50),
       (2,
        2,
        to_timestamp('2025-01-02 00:00:00', 'yyyy-mm-dd hh24:mi:ss'),
        10)
;

INSERT INTO shop.OrderProducts (order_id, product_id)
VALUES (1,
        1),
       (1,
        2),
       (2,
        1)
;

delete
from shop.Orders
where total_amount < 30;

select *
from shop.Users
where email like 'roma%';

select *
from shop.Products
where price > 30;

select *
from shop.Orders as o
         left join shop.Users u on o.user_id = u.id
where u.email like 'vasia%';

select u.*, SUM(o.total_amount)
from shop.Orders as o
         left join shop.Users u on o.user_id = u.id
group by u.id;

select u.*, AVG(p.price) as average_price
from shop.Users as u
         left join shop.Orders as o on u.id = o.user_id
         left join shop.OrderProducts as op on op.order_id = o.id
         left join shop.products p on p.id = OP.product_id
group by u.id