*Start Backend*:
- Go to `api/` folder.
- Run the following commands:
  - `docker compose up` - Docker download and pick up Redis.
  - `go mod tidy`
  - `go mod download`
  - `go run ./cmd/main.go`
  


*Start Frontend*:
- Go to `frontend/` folder.
- Run the following commands:
  - `npm install`
  - `npm start`
  