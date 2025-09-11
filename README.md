## Go REST API — Practice Project

This is a small REST API built in Go to practice core backend concepts: routing with Gin, persistence with SQLite, authentication with JWT, and password hashing with bcrypt.

### Tech
- **Gin**: HTTP router and middleware.
- **SQLite**: Simple file-based database (`api.db`).
- **golang-jwt/jwt/v5**: JWT generation and verification.
- **bcrypt**: Secure password hashing.

### Project Structure
```
main.go                 # Server entrypoint
database/db.go          # DB init and schema creation
middlewares/auth.go     # JWT auth middleware
models/                 # Data models (User, Event)
routes/                 # Handlers and route registration
utils/                  # JWT + password helpers
api.db                  # SQLite database (auto-created)
```

### Getting Started
1) Install Go 1.21+.
2) Install dependencies:
```
go mod download
```
3) Run the server:
```
go run .
```
The API listens on `http://localhost:8080` and creates `api.db` on first run.

### Authentication
- Sign up to create a user (password is hashed with bcrypt).
- Log in to receive a JWT.
- Include the token in the `Authorization` header for protected endpoints.

Note: In this practice project the token is sent as a raw string in `Authorization` (no `Bearer` prefix).

### Endpoints (all prefixed with `/api/v1`)

Public:
- `POST /api/v1/signup` — Create user
  - Body:
  ```json
  { "email": "you@example.com", "password": "yourPassword" }
  ```
  - 201 on success

- `POST /api/v1/login` — Get JWT
  - Body:
  ```json
  { "email": "you@example.com", "password": "yourPassword" }
  ```
  - 200 on success → `{ token: string }`

- `GET /api/v1/events` — List all events
- `GET /api/v1/events/:id` — Get a single event by id

Protected (requires `Authorization: <token>`):
- `POST /api/v1/events` — Create event
  - Body:
  ```json
  {
    "name": "Go Meetup",
    "description": "Learning Go",
    "location": "Online",
    "dateTime": "2025-09-10T18:00:00Z"
  }
  ```
- `PUT /api/v1/events/:id` — Update own event (name, description, location, dateTime)
- `DELETE /api/v1/events/:id` — Delete own event
- `POST /api/v1/events/:id/register` — Register current user to an event
- `DELETE /api/v1/events/:id/register` — Cancel current user's registration

### cURL Examples

Sign up:
```bash
curl -X POST http://localhost:8080/api/v1/signup \
  -H 'Content-Type: application/json' \
  -d '{"email":"you@example.com","password":"secret"}'
```

Log in (get token):
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"you@example.com","password":"secret"}' | jq -r .token)
echo "$TOKEN"
```

Create event:
```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: $TOKEN" \
  -d '{
    "name":"Go Meetup",
    "description":"Learning Go",
    "location":"Online",
    "dateTime":"2025-09-10T18:00:00Z"
  }'
```

List events:
```bash
curl http://localhost:8080/api/v1/events
```

Register for event:
```bash
curl -X POST http://localhost:8080/api/v1/events/1/register \
  -H "Authorization: $TOKEN"
```

Cancel registration:
```bash
curl -X DELETE http://localhost:8080/api/v1/events/1/register \
  -H "Authorization: $TOKEN"
```

### Notes and Limitations (since this is for practice)
- JWT secret is hard-coded and tokens are passed without the `Bearer` scheme.
- No advanced validation, pagination, or comprehensive error messages.
- Event ownership is enforced via `userId` embedded in JWT; only creators can update/delete their events.

### What I Practiced
- Setting up a Gin server and registering routes.
- Defining models and interacting with SQLite using `database/sql`.
- Hashing passwords with bcrypt and verifying during login.
- Generating and validating JWTs; applying an auth middleware.
- Implementing basic CRUD with ownership checks.

### Running Tests
No tests are included yet—this project focuses on hands-on API building. Potential next steps: add unit tests for utils, integration tests for routes, and input validation.


