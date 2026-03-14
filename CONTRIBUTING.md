# Contributing to file_renamer

Thank you for considering contributing! Here's how to get started.

## Reporting Issues

- Search [existing issues](../../issues) before opening a new one.
- Include your OS, Go version, and a minimal reproduction case.

## Submitting Changes

1. Fork the repository and create a branch from `master`:
   ```
   git checkout -b feat/your-feature
   ```
2. Make your changes, keeping commits focused and atomic.
3. Ensure the project builds and existing behaviour is preserved:
   ```
   go build ./...
   go vet ./...
   ```
4. Open a pull request against `master` with a clear description of *what* and *why*.

## Code Style

- Follow standard Go conventions (`gofmt`, `golint`).
- Keep functions small and focused.
- Avoid adding dependencies unless strictly necessary.

## Commit Messages

Use the conventional commits format:

```
<type>: <short summary>

Types: feat, fix, refactor, docs, chore, test
```

Example: `feat: add --dry-run flag`
