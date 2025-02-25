-- +migrate Up
ALTER TABLE users ADD COLUMN mobile VARCHAR(15) UNIQUE; -- 15 is the max length of a E164 mobile phone number

-- +migrate Down
ALTER TABLE users DROP COLUMN mobile;
