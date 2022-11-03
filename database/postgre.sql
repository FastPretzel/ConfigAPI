CREATE TABLE
  configs (
    config_id SERIAL PRIMARY KEY,
    service VARCHAR(50),
    config JSON,
    created_at TIMESTAMP
);
