CREATE TYPE dish_etag AS ENUM (
  'burgers',
  'sushi',
  'pizza',
  'desserts',
  'fast_food'
);

CREATE TABLE dish_tags (
  id serial PRIMARY KEY,
  dish_id int NOT NULL,
  tag dish_etag NOT NULL
);

ALTER TABLE dish_tags ADD FOREIGN KEY (dish_id) REFERENCES dishes (id);
