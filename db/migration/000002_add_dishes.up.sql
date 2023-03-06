CREATE TABLE dishes (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "rest_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE dishes ADD FOREIGN KEY (rest_id) REFERENCES restaurants (id);
