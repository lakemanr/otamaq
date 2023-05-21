CREATE TABLE users (
  id serial PRIMARY KEY,
  login varchar UNIQUE NOT NULL,
  full_name varchar NOT NULL,
  hashed_passwords varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE restaurants (
  id serial PRIMARY KEY,
  owner_id int NOT NULL,
  name varchar UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE clients (
  id serial PRIMARY KEY,
  owner_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE restaurants ADD CONSTRAINT restaurants_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE clients ADD CONSTRAINT clients_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE;