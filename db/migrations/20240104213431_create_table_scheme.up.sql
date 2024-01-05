CREATE TABLE schemes (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  created_at DATETIME,
  updated_at DATETIME,
  fund_house VARCHAR(255) NOT NULL,
  scheme_name VARCHAR(255) NOT NULL,
  scheme_type  VARCHAR(255) NOT NULL,
  scheme_category VARCHAR(255) NOT NULL
);
