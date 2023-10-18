-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE room (
  room_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  video_url VARCHAR(255) NOT NULL,
  playing BOOLEAN NOT NULL DEFAULT FALSE,
  progress FLOAT DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE room;
-- +goose StatementEnd
