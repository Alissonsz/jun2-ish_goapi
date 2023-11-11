-- room.GetById
-- Gets a room by its id
-- $1 roomId
SELECT
  room_id,
  name,
  video_url,
  playing,
  progress,
  created_at,
  updated_at
from room WHERE room_id = $1;
