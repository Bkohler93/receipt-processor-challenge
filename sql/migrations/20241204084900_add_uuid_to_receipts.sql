-- +goose Up
ALTER TABLE receipts ADD COLUMN uuid UUID NOT NULL DEFAULT gen_random_uuid();
UPDATE receipts SET uuid = gen_random_uuid();
ALTER TABLE receipts ADD CONSTRAINT unique_uuid UNIQUE (uuid);

-- +goose Down
ALTER TABLE receipts DROP CONSTRAINT unique_uuid;
ALTER TABLE receipts DROP COLUMN uuid;