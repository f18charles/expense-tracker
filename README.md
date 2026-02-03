# Expense Tracker

A lightweight, self-hosted expense tracking service written in Go. It provides a REST API and a small web UI to record, categorize, view and export personal or small-team expenses.

## Key features

- Create, read, update and delete expense records
- Categorization and tags
- Per-user authentication and session management (JWT or sessions)
- Monthly / custom-range summary and basic reporting
- CSV import / export
- Server-side rendered web UI and REST API for integrations
- Migrations and configuration-driven setup

## Project goals

- Simple, secure and maintainable codebase aimed at small deployments
- Clear separation between API, business logic and persistence layers
- Easy to run locally for development and in containers for production

## Non-goals

- Not intended to be a full accounting system; focus is on expense tracking and lightweight reporting
- Not providing complex tax or legal compliance features

## Tech stack

- Language: Go
- HTTP: net/http (or a light router such as gorilla/mux or chi)
- Persistence: SQL database (Postgres recommended) with migrations
- Frontend: minimal server-rendered templates and static assets
- Optional: Docker for containerized deployments

## Prerequisites

- Go (version 1.20+ recommended)
- Make (optional, many helper targets may be provided)
- A SQL database (Postgres is recommended) if running with persistence
- git

## Environment configuration

The application is configurable via environment variables. Common variables:

- PORT: HTTP port to listen on (default: 8080)
- DATABASE_URL: SQL connection string for the primary database
- JWT_SECRET: Secret key used to sign JWTs if JWT auth is enabled
- APP_ENV: app environment (development | staging | production)
- LOG_LEVEL: logging level (debug | info | warn | error)

Adjust or add variables as needed in your deployment environment. Do not commit secrets to source control.

## Running locally

1. Install dependencies and build:

   ```bash
   go build ./cmd/server
   ```

2. Set required environment variables (example):

   ```bash
   export PORT=8080
   export DATABASE_URL="postgres://user:pass@localhost:5432/expense_db?sslmode=disable"
   export JWT_SECRET="replace-with-a-secure-secret"
   ```

3. Run the server:

   ```bash
   ./server
   ```

For development convenience you may run with `go run`:

```bash
go run ./cmd/server
```

## Testing

- Unit and integration tests use `go test` across packages. Run all tests with:

  ```bash
  go test ./...
  ```

- Add tests for new features and keep them fast and deterministic. Use table-driven tests for clarity.

## Database migrations

Migrations are used to manage the database schema. The project contains a migrations directory and helpers to apply migrations in dev and production. Use your preferred migration tool and the migration files included in the repository (do not hardcode schema creation in your app at runtime).

## API and web UI (high level)

- The server exposes a RESTful JSON API for programmatic access and a small web UI for direct interaction.
- API roots and exact endpoints are defined in the codebase; check the handler and route definitions for details.
- Authentication protects write operations. Unauthenticated read-only endpoints may be available depending on configuration.

## Building and releasing

- Build binaries with `go build` or add CI pipelines to produce artifacts for your target platforms.
- Consider producing container images for deployment; a simple multi-stage Dockerfile is recommended for small binary images.

## Development notes

- Format code with `gofmt` / `go fmt` and run a linter (eg. `golangci-lint`) before submitting PRs.
- Keep business logic in internal packages and expose only handler/adaptor layers to the public packages.

## Contributing

Contributions are welcome. Please:

1. Fork the repository and create feature branches for non-trivial work
2. Run tests and linters locally before opening a pull request
3. Provide clear commit messages and a short PR description

## Security

- Treat secrets (database credentials, JWT secrets) securely. Use environment variables or a secrets manager in production.
- Report security issues privately instead of opening public issues. See the `SECURITY.md` if present for contact instructions.

## License

This project is licensed under the terms in `LICENSE.md`.

## Contact

For questions or support, open an issue or reach out to the repository owner.
