-- Room.CreateChatMessage
-- Persists a chat message in the database
-- $1 room_id
-- $2 author
-- $3 content
INSERT INTO chat_message 
  (room_id, author, content)
VALUES
  ($1, $2, $3)
RETURNING
  chat_message_id,
  room_id,
  author,
  content,
  created_at,
  updated_at;
