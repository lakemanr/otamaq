CREATE TABLE otamaq.dishes (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "rest_id" int,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE otamaq.dishes ADD FOREIGN KEY (rest_id) REFERENCES otamaq.restaurants (id);
