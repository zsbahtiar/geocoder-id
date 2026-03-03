# Installation

## CLI

Download the pre-built binary for your platform:

### macOS (Apple Silicon)

```bash
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-darwin-arm64 -o geocoder
chmod +x geocoder
sudo mv geocoder /usr/local/bin/
```

### Linux (x86_64)

```bash
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-linux-amd64 -o geocoder
chmod +x geocoder
sudo mv geocoder /usr/local/bin/
```

### Verify Installation

```bash
geocoder --version
```

## Go Library

### Using go get

```bash
go get github.com/zsbahtiar/geocoder-id
```

### Using go install (CLI only)

```bash
go install github.com/zsbahtiar/geocoder-id/cmd/geocoder@latest
```

## Database

The database (~233MB) is automatically downloaded on first use to `~/.geocoder-id/geocoder.duckdb`.

### Custom Database Location

Set the `GEOCODER_DB_PATH` environment variable:

```bash
export GEOCODER_DB_PATH=/path/to/geocoder.duckdb
```

### Database Search Order

1. `$GEOCODER_DB_PATH` environment variable
2. `./geocoder.duckdb` (current directory)
3. `./data/geocoder.duckdb`
4. `<binary-dir>/geocoder.duckdb`
5. `<binary-dir>/data/geocoder.duckdb`
6. `~/.geocoder-id/geocoder.duckdb` (auto-downloaded)
