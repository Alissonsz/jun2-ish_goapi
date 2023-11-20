-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE chat_message (
  chat_message_id SERIAL PRIMARY KEY,
  author VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  room_id INTEGER NOT NULL,
  CONSTRAINT FK_chat_message_room FOREIGN KEY (room_id) REFERENCES room (room_id),
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE chat_message;
-- +goose StatementEnd
