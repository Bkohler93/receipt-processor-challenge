-- +goose Up
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    receipt_id INT NOT NULL,
    short_description TEXT NOT NULL,
    price DECIMAL(12,2) NOT NULL,
    CONSTRAINT fk_receipt FOREIGN KEY (receipt_id) REFERENCES receipts (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE items;