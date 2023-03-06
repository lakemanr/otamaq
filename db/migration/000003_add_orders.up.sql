CREATE TABLE orders (
  id serial PRIMARY KEY,
  client_id int NOT NULL,
  rest_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE order_items (
  id serial PRIMARY KEY,
  order_id int NOT NULL,
  dish_id int NOT NULL,
  quantity int NOT NULL DEFAULT 1
);

ALTER TABLE orders ADD FOREIGN KEY (client_id) REFERENCES clients (id);

ALTER TABLE orders ADD FOREIGN KEY (rest_id) REFERENCES restaurants (id);

ALTER TABLE order_items ADD FOREIGN KEY (order_id) REFERENCES orders (id);

ALTER TABLE order_items ADD FOREIGN KEY (dish_id) REFERENCES dishes (id);
