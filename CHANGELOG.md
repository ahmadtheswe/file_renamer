# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [Unreleased]

### Added
- `cmd/root.go` — command logic moved out of `main.go` for a cleaner entry point.
- Mandatory flag validation: `-dir`, `-ext`, and `-prefix` are now required; missing flags print usage and exit with a non-zero code.
- Custom `-h`/`-help` output with a structured usage message.

---

## [0.1.0] - 2026-03-14

### Added
- Initial release.
- CLI flags: `--dir`, `--ext`, `--prefix`, `--v`.
- Bulk rename files in a directory by adding a numbered, zero-padded prefix.
- Skips files that already carry the target prefix.
- Confirmation prompt before any renaming is performed.
