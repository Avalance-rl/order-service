CREATE TYPE order_status AS ENUM
(
    'UNPAID',
    'PAID',
    'COMPLETED'
);

CREATE TABLE IF NOT EXISTS orders
(
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    order_status order_status NOT NULL,
    product_list VARCHAR(64)[] NOT NULL,
    total_price INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);