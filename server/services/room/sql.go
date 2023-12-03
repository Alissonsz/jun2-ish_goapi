package room

import _ "embed"

var (
	//go:embed sql/insert.sql
	insertQuery string
	//go:embed sql/get_by_id.sql
	getByIdQuery string
	//go:embed sql/insert_chat_message.sql
	insertChatMessageQuery string
	//go:embed sql/get_chat_messages.sql
	getChatMessagesQuery string
	//go:embed sql/update.sql
	updateQuery string
	//go:embed sql/insert_playlist_item.sql
	insertPlaylistItemQuery string
	//go:embed sql/get_playlist_items.sql
	getPlaylistItemsQuery string
)
