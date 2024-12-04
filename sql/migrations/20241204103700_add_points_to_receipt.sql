-- +goose Up
ALTER TABLE receipts ADD COLUMN points INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE receipts DROP COLUMN points;