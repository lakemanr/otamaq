CREATE SCHEMA otamaq;

CREATE TABLE otamaq.restaurants (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE otamaq.clients (
  "id" serial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "login" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);