CREATE TABLE IF NOT EXISTS sites (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name VARCHAR (255) NOT NULL,
  url VARCHAR (255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
