# case-study-dining2

Go live-coding workspace. Standard library only — no external dependencies, so it runs without network access.

## Requirements

- Go 1.24+

## Quick start

```bash
make run          # starts the API on :8080 (override with HTTP_ADDR)
curl localhost:8080/healthz
make test         # go test ./...
make check        # gofmt + go vet + go test -race
```

## Layout

```
cmd/api/            entrypoint: config, server timeouts, graceful shutdown
internal/server/    router (Go 1.22+ method/path patterns) + middleware chain
internal/httpx/     JSON read/write helpers, error responses, middleware
internal/apperr/    domain errors (ErrNotFound, ErrValidation, ErrConflict) -> HTTP status
```

## Conventions

- Handlers decode with `httpx.ReadJSON` (unknown fields rejected, 1 MiB body limit).
- Domain code returns wrapped `apperr` errors (`fmt.Errorf("%w: ...", apperr.ErrValidation)`);
  handlers call `httpx.WriteError`, which maps them to status codes and hides internal errors.
- Middleware: panic recovery, `X-Request-ID`, structured request logging (`log/slog`).
- New feature = `internal/<feature>/` with model, store interface + in-memory implementation
  (guarded by `sync.RWMutex`), service, handler, table-driven tests.
