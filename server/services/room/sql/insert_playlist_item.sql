-- room.CreatePlaylistItem
-- Persists a playlist item struct in the database
-- $1 room_id
-- $2 video_url
-- $3 name

INSERT INTO playlist_item
  (room_id, video_url, name) 
VALUES
  ($1, $2, $3)
RETURNING
  playlist_item_id,
  room_id,
  video_url,
  name,
  created_at,
  updated_at,
  deleted_at;
