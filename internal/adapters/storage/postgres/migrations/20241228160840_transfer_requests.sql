-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transfer_requests (
  id          VARCHAR(36) PRIMARY KEY,
  item_id    VARCHAR(36) NOT NULL,
  quantity  INT NOT NULL,
  from_inventory_id VARCHAR(36) NOT NULL,
  to_inventory_id VARCHAR(36) NOT NULL,
  from_section_id VARCHAR(36),
  to_section_id VARCHAR(36),
  transfer_status VARCHAR(255) NOT NULL,
  updated_at  TIMESTAMPTZ DEFAULT now(),
  created_at  TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transfer_requests;
-- +goose StatementEnd
