# Geocoder ID

Offline reverse geocoding for Indonesia. Convert coordinates (latitude/longitude) to administrative addresses down to village/sub-district level.

## Features

- **Offline** - no internet required (after database is downloaded)
- **Fast** - spatial index for O(log n) queries
- **Auto-download** - database automatically downloaded on first use
- **Dual-use** - works as CLI tool or Go library

## Installation

### CLI

Download binary from [Releases](https://github.com/zsbahtiar/geocoder-id/releases):

```bash
# macOS (Apple Silicon)
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-darwin-arm64 -o geocoder
chmod +x geocoder

# macOS (Intel)
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-darwin-amd64 -o geocoder
chmod +x geocoder

# Linux
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-linux-amd64 -o geocoder
chmod +x geocoder
```

### Go Library

```bash
go get github.com/zsbahtiar/geocoder-id
```

## Usage

### CLI

```bash
# Reverse geocoding
geocoder geocode --coords="-6.2088 106.8456"

# Multiple coordinates (semicolon separated)
geocoder geocode --coords="-6.2088 106.8456;-2.04963 110.18774"

# JSON output
geocoder geocode --coords="-6.2088 106.8456" --output=json

# Specific level (province, regency, district, village)
geocoder geocode --coords="-6.2088 106.8456" --level=district
```

### Go Library

```go
package main

import (
    "fmt"
    "log"

    geocoder "github.com/zsbahtiar/geocoder-id"
)

func main() {
    // Auto-downloads database if not found
    gc, err := geocoder.NewDefault()
    if err != nil {
        log.Fatal(err)
    }
    defer gc.Close()

    // Reverse geocode
    result := gc.ReverseGeocode(-6.2088, 106.8456)

    if result.Village != nil {
        fmt.Printf("Village: %s (%s)\n", result.Village.Name, result.Village.Code)
    }
}
```

## Output

**Table** (default):
```
| province_code | province_name | regency_code | regency_name                      | district_code | district_name | village_code  | village_name  |
|---------------|---------------|--------------|-----------------------------------|---------------|---------------|---------------|---------------|
| 31            | DKI Jakarta   | 31.74        | Kota Administrasi Jakarta Selatan | 31.74.02      | Setiabudi     | 31.74.02.1006 | Pasar Manggis |
```

**JSON** (`--output=json`):
```json
[
  {
    "province": { "code": "31", "name": "DKI Jakarta" },
    "regency": { "code": "31.74", "name": "Kota Administrasi Jakarta Selatan" },
    "district": { "code": "31.74.02", "name": "Setiabudi" },
    "village": { "code": "31.74.02.1006", "name": "Pasar Manggis" }
  }
]
```

## Build from Source

```bash
git clone https://github.com/zsbahtiar/geocoder-id
cd geocoder-id

# Build
make build

# Run
./geocoder geocode --coords="-6.2088 106.8456"
```

## License

MIT
