CREATE TABLE
  mf_api_relations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW() ON UPDATE NOW(),
    relation_id BIGINT,
    scheme_id BIGINT,
    FOREIGN KEY (scheme_id) REFERENCES schemes (id)
  );