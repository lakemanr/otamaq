ALTER TABLE IF EXISTS clients DROP CONSTRAINT IF EXISTS clients_owner_login_fkey;
ALTER TABLE IF EXISTS restaurants DROP CONSTRAINT IF EXISTS restaurants_owner_login_fkey;
DROP TABLE IF EXISTS restaurants;
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS users;