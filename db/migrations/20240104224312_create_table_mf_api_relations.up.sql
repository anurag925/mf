CREATE TABLE mf_api_relations (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  created_at DATETIME,
  updated_at DATETIME,
  relation_id BIGINT,
  scheme_id BIGINT,

  FOREIGN KEY (scheme_id) REFERENCES schemes(id)
);
