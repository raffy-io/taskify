-- +goose Up
CREATE TABLE homepage (
  id    BIGSERIAL PRIMARY KEY,
  image TEXT      NOT NULL
);

-- +goose Down
DROP TABLE homepage;