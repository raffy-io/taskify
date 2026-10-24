-- +goose Up
CREATE TABLE homepage (
  id    BIGSERIAL PRIMARY KEY,
  hero_image TEXT NOT NULL,
  hero_image_alt TEXT NOT NULL,
  hero_title TEXT NOT NULL,
  hero_text TEXT NOT NULL
);

-- +goose Down
DROP TABLE homepage;