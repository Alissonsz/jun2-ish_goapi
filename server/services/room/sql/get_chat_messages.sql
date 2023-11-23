-- room.GetChatMessages
-- Get room's chat messages by room id
-- $1 roomId
SELECT
  author,
  content,
  room_id,
  chat_message_id,
  created_at,
  updated_at,
  deleted_at
from chat_message WHERE room_id = $1;
