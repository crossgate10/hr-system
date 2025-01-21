CREATE TABLE employees
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(64) NOT NULL,
    title        VARCHAR(64),
    role         VARCHAR(64),
    contact_info VARCHAR(64),
    salary       FLOAT,
    status       VARCHAR(10),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
