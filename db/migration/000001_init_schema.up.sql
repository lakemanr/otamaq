CREATE TABLE restaurants (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE clients (
  "id" serial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "login" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);