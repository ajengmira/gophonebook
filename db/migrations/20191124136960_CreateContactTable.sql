-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE contact (
  id            INT UNSIGNED         NOT NULL AUTO_INCREMENT PRIMARY KEY,
  contact_name  VARCHAR(150)         NOT NULL,
  phone_number  VARCHAR(150)         NOT NULL,
  user_id       INT(10)              NOT NULL,
  created_at    DATETIME,
  updated_at    DATETIME,
  FOREIGN KEY (user_id) REFERENCES user (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE contact;
