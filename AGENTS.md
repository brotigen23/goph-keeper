# Repository Guidelines

## Project Structure

- `backend/services/auth/` contains the Go authentication service. Its entry point is `cmd/main.go`, and `internal/app/` assembles and starts the Gin server.
- Keep entities and repository interfaces in `internal/domain/`, business logic in `internal/usecase/`, and repository implementations in `internal/infrastructure/repository/`. HTTP routing, request and response types, and JWT handling live in `internal/transport/http/`.
- `backend/services/shared/pkg/` contains reusable cryptography, middleware, and utility packages. See `docs/Specifications.md` and `docs/backend/auth/` for requirements and auth design diagrams.

## Build, Test, and Run

Use Go 1.23.7 or newer. The root Makefile references legacy client and server directories, so run commands from the service module directories. For local development, initialize a workspace from the repository root with `go work init ./backend/services/auth ./backend/services/shared` if you do not already have one. Keep that workspace local.

- In `backend/services/auth/`, run `go run github.com/google/wire/cmd/wire ./internal/transport/http` to generate the auth injector on a fresh checkout or after changing its providers. The generated `wire_gen.go` is ignored by Git.
- In the auth module, run `go build -o bin/auth ./cmd` to build the service or `make run` to start it on port 8080.
- In each module directory, run `go test ./...` to check all packages in that module. Use `go test -cover ./...` to inspect coverage.

## Coding Style

Run `gofmt -w` on changed Go files. Use tabs for indentation, short lowercase package names, and mixed-case Go identifiers. Follow existing descriptive filenames such as `create_user.go`. Keep service-specific Gin types in the HTTP transport layer.

## Testing

Use Go’s standard `testing` package. Place automated tests next to the code in `*_test.go` files and name test functions `TestXxx`. The repository has no defined coverage threshold. For a manual auth check, start the service and send the requests in `backend/services/auth/test/` with matching, nonempty test credentials, registering before logging in. The in-memory repository resets on restart.

## Commits and Pull Requests

Use short, descriptive commit subjects. History commonly uses prefixes such as `feat:`, `test:`, and `doc:`. In pull requests, summarize the change, link related issues when available, and list verification commands or manual checks. Describe any API behavior changes with example requests when helpful.
