# file_renamer

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)

A simple CLI tool to bulk rename files in a directory by adding a numbered prefix.

## Usage

```bash
go run main.go [flags]
```

Or build first:

```bash
go build -o file_renamer .
./file_renamer [flags]
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--dir` | `.` | Path to the target directory |
| `--ext` | *(none)* | Filter files by extension (e.g. `.jpg`, `.txt`) |
| `--prefix` | *(none)* | Prefix to prepend to renamed files |
| `--v` | `false` | Verbose output — prints each rename operation |
| `--help` | *(none)* | Show usage information |

## How it works

1. Scans the directory (`--dir`) for files matching the given extension (`--ext`).
2. Skips files that already start with the given prefix.
3. Shows a summary and prompts for confirmation before making any changes.
4. Renames matching files to `<prefix>_<index><ext>`, where the index is zero-padded based on the total file count. The index starts after any already-prefixed files so existing numbered files are not overwritten.

## Examples

Rename all `.jpg` files in a folder with the prefix `photo`:

```bash
go run main.go --dir C:\Users\me\Pictures --ext .jpg --prefix photo
```

Same, with verbose output to see each rename:

```bash
go run main.go --dir C:\Users\me\Pictures --ext .jpg --prefix photo --v
```

Rename all files (no extension filter) in the current directory:

```bash
go run main.go --prefix backup
```

### Example output

```
Contents of directory 'C:\Users\me\Pictures':
Total files with extension '.jpg' that will be renamed: 3
Are you sure you want to rename these files? (y/n): y
Renaming 'IMG001.jpg' to 'photo_1.jpg'
Renaming 'IMG002.jpg' to 'photo_2.jpg'
Renaming 'IMG003.jpg' to 'photo_3.jpg'
Renaming completed.
```

## Build and run with Go 1.20 or later:

### Build and run:
```bash
go build -o file_renamer .
./file_renamer --dir C:\path\to\directory --ext .txt --prefix myfile
```

### Run without building:
```bash
go run main.go --dir C:\path\to\directory --ext .txt --prefix myfile
```

### Makefile commands
#### Build with Makefile
```bash
make build
```
#### Run with Makefile
```bash
make run
```
#### Clean build artifacts
```bash
make clean
```

## Contributing

Contributions are welcome! Please read the [Contributing Guide](CONTRIBUTING.md) for details on branching workflow, code style, and commit message conventions.

## Code of Conduct

This project follows the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold these standards.

## Changelog

All notable changes are documented in the [Changelog](CHANGELOG.md).

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
