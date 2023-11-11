-- room.Create
-- Persists a room struct in the database
-- $1 name
-- $2 video_url
-- $3 playing
-- $4 progress
INSERT INTO room
  (name, video_url, playing, progress)
VALUES
  ($1, $2, $3, $4)
RETURNING 
  room_id,
  name,
  video_url,
  playing,
  progress,
  created_at,
  updated_at;
