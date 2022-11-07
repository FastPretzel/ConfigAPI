CREATE TABLE service (
    service_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE
);

CREATE TABLE
  config (
    config_id SERIAL PRIMARY KEY,
    service_id INT,
    config JSON,
    version INT,
    in_use BOOL DEFAULT FALSE,
    created_at TIMESTAMP,
    FOREIGN KEY (service_id) REFERENCES service (service_id) ON DELETE CASCADE
);
