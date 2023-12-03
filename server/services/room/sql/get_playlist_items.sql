-- room.GetPlaylistItems
-- Get room's playlist items by room id
-- $1 roomId

SELECT 
  playlist_item_id, 
  room_id,
  video_url,
  name,
  created_at,
  updated_at,
  deleted_at
from playlist_item
WHERE room_id = $1;
