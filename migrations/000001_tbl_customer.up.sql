CREATE TABLE IF NOT EXISTS tbl_customer(
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255),
    balance DECIMAL(10, 2) DEFAULT 0.00,
    created_at TIMESTAMP WITH TIME ZONE  DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE  DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);