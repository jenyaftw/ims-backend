-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventory_item (
  id          VARCHAR(36) PRIMARY KEY,
  name        VARCHAR(255) NOT NULL,
  description TEXT,
  quantity   INT NOT NULL,
  sku       VARCHAR(255) NOT NULL,
  inventory_id VARCHAR(36) NOT NULL,
  section_id VARCHAR(36),
  updated_at  TIMESTAMPTZ DEFAULT now(),
  created_at  TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory_item;
-- +goose StatementEnd
