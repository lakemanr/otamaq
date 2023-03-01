CREATE TYPE otamaq.dish_tag AS ENUM (
  'burgers',
  'sushi',
  'pizza',
  'desserts',
  'fast_food'
);

CREATE TABLE otamaq.dish_tags (
  id serial PRIMARY KEY,
  dish_id int,
  tag otamaq.dish_tag
);

ALTER TABLE otamaq.dish_tags ADD FOREIGN KEY (dish_id) REFERENCES otamaq.dishes (id);
