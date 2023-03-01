CREATE TABLE otamaq.orders (
  id serial PRIMARY KEY,
  client_id int,
  rest_id int,
  created_at timestamp DEFAULT (now())
);

CREATE TABLE otamaq.order_items (
  id serial PRIMARY KEY,
  order_id int,
  dish_id int,
  quantity int DEFAULT 1
);

ALTER TABLE otamaq.orders ADD FOREIGN KEY (client_id) REFERENCES otamaq.clients (id);

ALTER TABLE otamaq.orders ADD FOREIGN KEY (rest_id) REFERENCES otamaq.restaurants (id);

ALTER TABLE otamaq.order_items ADD FOREIGN KEY (order_id) REFERENCES otamaq.orders (id);

ALTER TABLE otamaq.order_items ADD FOREIGN KEY (dish_id) REFERENCES otamaq.dishes (id);
