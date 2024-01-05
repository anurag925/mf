CREATE TABLE
  navs (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT (NOW()),
    updated_at DATETIME DEFAULT NOW() ON UPDATE NOW(),
    date DATETIME,
    value DECIMAL(13, 5),
    scheme_id BIGINT,

    FOREIGN KEY (scheme_id) REFERENCES schemes (id),
    INDEX(date, value)
  );