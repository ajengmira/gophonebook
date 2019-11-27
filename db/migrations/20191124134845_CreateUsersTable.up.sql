CREATE TABLE users (
  id         INT UNSIGNED         NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username   VARCHAR(150)         NOT NULL,
  email      VARCHAR(150)         NOT NULL,
  password   VARCHAR(150)         NOT NULL,
  created_at DATETIME,
  updated_at DATETIME
);
