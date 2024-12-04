-- +goose Up
CREATE TABLE receipts (
    id SERIAL PRIMARY KEY,
    retailer TEXT NOT NULL,
    purchase_date DATE NOT NULL,
    purchase_time TIME NOT NULL,
    total DECIMAL(12,2) NOT NULL
);

-- +goose Down
DROP TABLE receipts;