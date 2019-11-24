-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id         INT UNSIGNED         NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username   VARCHAR(150)         NOT NULL,
  email      VARCHAR(150)         NOT NULL,
  password   VARCHAR(150)         NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
