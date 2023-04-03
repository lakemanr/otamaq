CREATE TABLE dishes (
  id serial PRIMARY KEY,
  name varchar NOT NULL,
  rest_id int NOT NULL,
  quantity int NOT NULL DEFAULT 0 CHECK (quantity >= 0), 
  created_at timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE dishes ADD FOREIGN KEY (rest_id) REFERENCES restaurants (id);
