-- room.DeletePlaylistItem
-- $1 playlist_item_id

UPDATE playlist_item
SET
  deleted_at = now()
WHERE playlist_item_id = $1
RETURNING 
  playlist_item_id,
  room_id,
  name,
  video_url,
  created_at,
  updated_at,
  deleted_at;