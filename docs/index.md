---
layout: home

hero:
  name: Geocoder ID
  text: Offline Reverse Geocoding
  tagline: Convert coordinates to Indonesian administrative addresses - offline, fast, and accurate
  actions:
    - theme: brand
      text: Get Started
      link: /guide/
    - theme: alt
      text: View on GitHub
      link: https://github.com/zsbahtiar/geocoder-id

features:
  - icon: 🔌
    title: Offline First
    details: No internet required after initial database download. Works completely offline.
  - icon: ⚡
    title: Fast Queries
    details: Spatial index for O(log n) queries. Get results in milliseconds.
  - icon: 📦
    title: Dual Use
    details: Works as CLI tool or importable Go library. One codebase, multiple uses.
  - icon: 🗺️
    title: Complete Coverage
    details: 38 provinces, 514 regencies, 7,277 districts, and 83,288 villages.
---

## Quick Install

### CLI

```bash
# macOS (Apple Silicon)
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-darwin-arm64 -o geocoder
chmod +x geocoder

# Linux
curl -L https://github.com/zsbahtiar/geocoder-id/releases/latest/download/geocoder-linux-amd64 -o geocoder
chmod +x geocoder
```

### Go Library

```bash
go get github.com/zsbahtiar/geocoder-id
```

## Quick Example

```bash
$ geocoder geocode --coords="-6.2088 106.8456"

| province_code | province_name | regency_code | regency_name                      | district_code | district_name | village_code  | village_name  |
|---------------|---------------|--------------|-----------------------------------|---------------|---------------|---------------|---------------|
| 31            | DKI Jakarta   | 31.74        | Kota Administrasi Jakarta Selatan | 31.74.02      | Setiabudi     | 31.74.02.1006 | Pasar Manggis |
```
