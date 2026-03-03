# CLI Reference

## Commands

### geocode

Reverse geocode coordinates to administrative addresses.

```bash
geocoder geocode [flags]
```

#### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--coords` | string | required | Coordinates in "lat lon" format. Multiple coords separated by `;` |
| `--output` | string | `table` | Output format: `table` or `json` |
| `--level` | string | `village` | Query level: `province`, `regency`, `district`, `village` |
| `--db` | string | auto | Path to database file |
| `--debug` | bool | `false` | Enable debug logging |

#### Examples

```bash
# Single coordinate
geocoder geocode --coords="-6.2088 106.8456"

# Multiple coordinates
geocoder geocode --coords="-6.2088 106.8456;-7.7956 110.3695"

# JSON output
geocoder geocode --coords="-6.2088 106.8456" --output=json

# District level only
geocoder geocode --coords="-6.2088 106.8456" --level=district

# Custom database path
geocoder geocode --coords="-6.2088 106.8456" --db=/path/to/geocoder.duckdb

# Debug mode
geocoder geocode --coords="-6.2088 106.8456" --debug
```

### version

Show version information.

```bash
geocoder version
```

## Global Flags

| Flag | Description |
|------|-------------|
| `--help` | Help for any command |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `GEOCODER_DB_PATH` | Custom path to database file |

## Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Error (invalid input, database not found, etc.) |
