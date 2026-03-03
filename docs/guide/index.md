# Introduction

Geocoder ID is an offline reverse geocoding tool for Indonesia. It converts latitude/longitude coordinates to administrative addresses down to the village (kelurahan/desa) level.

## What is Reverse Geocoding?

Reverse geocoding is the process of converting geographic coordinates (latitude and longitude) into a human-readable address. For example:

- **Input**: `-6.2088, 106.8456`
- **Output**: Pasar Manggis, Setiabudi, Jakarta Selatan, DKI Jakarta

## Features

### Offline First
Once the database is downloaded (~233MB), no internet connection is required. Perfect for:
- Embedded systems
- Air-gapped environments
- Batch processing without API rate limits

### Fast Performance
Uses DuckDB with spatial indexing for O(log n) query performance. Typical queries complete in under 10ms.

### Complete Coverage
Covers all administrative levels in Indonesia:

| Level | Count |
|-------|-------|
| Province (Provinsi) | 38 |
| Regency (Kabupaten/Kota) | 514 |
| District (Kecamatan) | 7,277 |
| Village (Kelurahan/Desa) | 83,288 |

### Dual Use
- **CLI**: Download binary and run from terminal
- **Go Library**: Import and use in your Go applications

## Next Steps

- [Installation](/guide/installation) - Install the CLI or Go library
- [Quick Start](/guide/quick-start) - Get up and running in minutes
