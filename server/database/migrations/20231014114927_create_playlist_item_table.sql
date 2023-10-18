-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE playlist_item (
  playlist_item_id SERIAL PRIMARY KEY,
  room_id INTEGER NOT NULL,
  video_url VARCHAR(255) NOT NULL,
  name VARCHAR(255),
  CONSTRAINT FK_playlist_item_room FOREIGN KEY (room_id) REFERENCES room (room_id),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE playlist_item;
-- +goose StatementEnd
