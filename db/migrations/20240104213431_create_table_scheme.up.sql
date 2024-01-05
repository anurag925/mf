CREATE TABLE
  schemes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW() ON UPDATE NOW(),
    fund_house VARCHAR(255) NOT NULL,
    scheme_name VARCHAR(255) NOT NULL,
    scheme_type VARCHAR(255) NOT NULL,
    scheme_category VARCHAR(255) NOT NULL
  );