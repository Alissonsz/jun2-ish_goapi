-- room.Update
-- Persists a room struct in the database
-- $1 room_id
-- $2 name
-- $3 video_url
-- $4 playing
-- $5 progress
UPDATE room
SET
  name = $2,
  video_url = $3,
  playing = $4,
  progress = $5,
  updated_at = now()
WHERE room_id = $1
RETURNING 
  room_id,
  name,
  video_url,
  playing,
  progress,
  created_at,
  updated_at;
