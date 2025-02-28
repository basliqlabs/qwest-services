-- +migrate Up
CREATE TABLE users(
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(32) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password_hash VARCHAR(64) NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);
-- +migrate Down
DROP TABLE IF EXISTS users;