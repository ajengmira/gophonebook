CREATE TABLE contact (
  id            INT UNSIGNED         NOT NULL AUTO_INCREMENT PRIMARY KEY,
  contact_name  VARCHAR(150)         NOT NULL,
  phone_number  VARCHAR(150)         NOT NULL,
  created_at    DATETIME,
  updated_at    DATETIME
);
