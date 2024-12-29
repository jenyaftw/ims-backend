-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventory_sections (
  id          VARCHAR(36) PRIMARY KEY,
  inventory_id VARCHAR(36) NOT NULL,
  name        VARCHAR(255) NOT NULL,
  description TEXT,
  updated_at  TIMESTAMPTZ DEFAULT now(),
  created_at  TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventories;
-- +goose StatementEnd
