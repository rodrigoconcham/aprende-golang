# Copilot instructions for aprende-golang

Purpose: Give an AI coding agent the minimal, actionable knowledge to be productive in this repository.

Quick commands
- Build/run: `go run .` (from repo root)
- Build binary: `go build ./...`
- Tests: `go test ./...` or `go test ./mathutil`
- Module info: `cat go.mod` shows `module aprende-golang` and `go 1.25.4`

Project layout (key files)
- `main.go` — program entrypoint, demonstrates concurrent HTTP fetches
- `mathutil/mathutil.go` — small package with exported functions `Add`, `Resta` and exported var `Edad`
- `go.mod` — module and Go version

Repository-specific patterns & notes
- Concurrency: `main.go` launches goroutines to fetch remote URLs. The code currently calls `wg.Add(1)` but the goroutine does not call `wg.Done()` nor does `main` call `wg.Wait()`; any concurrency fixes must add `wg.Done()` in the goroutine and call `wg.Wait()` before finishing `main`.
  - Example change request: "Fix `sync.WaitGroup` usage in `main.go`: pass the `wg` to each goroutine (or use an inline closure), call `defer wg.Done()` inside the goroutine, and `wg.Wait()` before printing total duration."

- Package exports: `mathutil` uses exported identifiers (`Add`, `Resta`, `Edad`) — follow this pattern when adding helpers that should be accessible from `main`.

- Language and messages: log/output strings are Spanish (for example `Fetched %s en %v - Status %s` and `Tiempo total de la secuencia`). When adding user-visible messages, prefer Spanish for consistency.

- External dependency: `main.go` calls the Rick and Morty API (`https://rickandmortyapi.com`). Network requests are expected to run in CI or dev machines with Internet access; consider mocking or using a test HTTP server when writing unit tests.

Tests and adding coverage
- No tests exist currently. When adding tests:
  - Put unit tests next to packages (`mathutil/mathutil_test.go`) and use `go test ./...`.
  - Prefer table-driven tests for functions like `Add`/`Resta`.
  - For `main.go` behavior that depends on network I/O, abstract the HTTP call into a function or interface so tests can inject a mock client.

Editing guidelines (concrete)
- Small PRs only: make single-purpose changes (e.g., "fix WaitGroup", "add tests for mathutil.Add").
- Preserve Spanish logging text unless the change is explicitly about localization.
- When touching concurrency in `main.go`, ensure `wg.Done()` and `wg.Wait()` are used correctly and measure elapsed time only after `wg.Wait()`.

Examples of useful prompts for the agent
- "Create `mathutil/mathutil_test.go` with table-driven tests for `Add` and `Resta` and run `go test ./mathutil`."
- "Fix `main.go` concurrency: ensure each goroutine calls `defer wg.Done()` and `main` waits on `wg.Wait()` before printing total time."
- "Refactor `fetchURL` to accept an `http.Client` parameter for easier testing; update `main.go` to pass `http.DefaultClient`."

Where to look first
- Start with `main.go` to understand the app flow and concurrency usage.
- Inspect `mathutil/mathutil.go` for small helper patterns and exported API.
- Use `go.mod` to confirm the expected Go toolchain version.

If anything is unclear
- Ask for which area to prioritize (concurrency fixes, tests, or refactors). Include desired language (Spanish vs English) for new user-visible text.

End of file
