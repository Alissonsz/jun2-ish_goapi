Using this as boilerplate: https://github.com/nixsolutions/golang-gin-boilerplate

### Utilities
- Create migration: `make create-migration name=${name}`
- Run migrations: `make run-migrations`
- Create db container using docker: `make create-docker-db`

### Types
```
interface ChatMessage {
  author: string;
  content: string;
}

interface PlaylistItem {
  id: number;
  name: string;
  url: string;
}

export interface Room {
  id: string;
  name: string;
  userCount: number;
  videoUrl: string;
  messages: ChatMessage[];
  playing: boolean;
  progress: number;
  playlist: PlaylistItem[];
}
```

### REST Routes
- GET `/rooms`
  - params: `null`
  - response: `{ rooms: []Room }`
- POST `/room`
  - params: `{ name: string; videoUrl?: string }`
  - response: `{ room: Room }`
- GET `/room/{id}`
  - response: `{ room: Room }` or `404` if not found    
