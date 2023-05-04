CREATE TABLE users (
  login varchar PRIMARY KEY,
  full_name varchar NOT NULL,
  hashed_passwords varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE restaurants (
  id serial PRIMARY KEY,
  owner_login varchar NOT NULL,
  name varchar UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE clients (
  id serial PRIMARY KEY,
  owner_login varchar UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE restaurants ADD CONSTRAINT restaurants_owner_login_fkey FOREIGN KEY (owner_login) REFERENCES users (login) ON DELETE CASCADE;
ALTER TABLE clients ADD CONSTRAINT clients_owner_login_fkey FOREIGN KEY (owner_login) REFERENCES users (login) ON DELETE CASCADE;