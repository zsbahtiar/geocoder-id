# Quick Start

## CLI Usage

### Basic Geocoding

```bash
geocoder geocode --coords="-6.2088 106.8456"
```

Output:
```
| province_code | province_name | regency_code | regency_name                      | district_code | district_name | village_code  | village_name  |
|---------------|---------------|--------------|-----------------------------------|---------------|---------------|---------------|---------------|
| 31            | DKI Jakarta   | 31.74        | Kota Administrasi Jakarta Selatan | 31.74.02      | Setiabudi     | 31.74.02.1006 | Pasar Manggis |
```

### Multiple Coordinates

Separate coordinates with semicolons:

```bash
geocoder geocode --coords="-6.2088 106.8456;-7.7956 110.3695"
```

Output:
```
| province_code | province_name                 | regency_code | regency_name                      | district_code | district_name | village_code  | village_name    |
|---------------|-------------------------------|--------------|-----------------------------------|---------------|---------------|---------------|-----------------|
| 31            | Daerah Khusus Ibukota Jakarta | 31.74        | Kota Administrasi Jakarta Selatan | 31.74.02      | Setiabudi     | 31.74.02.1006 | Pasar Manggis   |
| 34            | Daerah Istimewa Yogyakarta    | 34.04        | Kabupaten Sleman                  | 34.04.05      | Depok         | 34.04.05.2003 | Caturtunggal    |
```

### JSON Output

```bash
geocoder geocode --coords="-6.2088 106.8456;-7.7956 110.3695" --output=json
```

```json
[
  {
    "province": { "code": "31", "name": "Daerah Khusus Ibukota Jakarta" },
    "regency": { "code": "31.74", "name": "Kota Administrasi Jakarta Selatan" },
    "district": { "code": "31.74.02", "name": "Setiabudi" },
    "village": { "code": "31.74.02.1006", "name": "Pasar Manggis" }
  },
  {
    "province": { "code": "34", "name": "Daerah Istimewa Yogyakarta" },
    "regency": { "code": "34.04", "name": "Kabupaten Sleman" },
    "district": { "code": "34.04.05", "name": "Depok" },
    "village": { "code": "34.04.05.2003", "name": "Caturtunggal" }
  }
]
```

### Query Specific Level

```bash
# Province only
geocoder geocode --coords="-6.2088 106.8456" --level=province

# District level
geocoder geocode --coords="-6.2088 106.8456" --level=district
```

## Go Library Usage

### Basic Example

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

### Batch Processing

```go
coords := []geocoder.Coord{
    {Lat: -6.2088, Lon: 106.8456},
    {Lat: -2.04963, Lon: 110.18774},
    {Lat: -7.7956, Lon: 110.3695},
}

for _, c := range coords {
    result := gc.ReverseGeocode(c.Lat, c.Lon)
    fmt.Printf("%.4f, %.4f -> %s\n", c.Lat, c.Lon, result.Village.Name)
}
```

### Query Specific Level

```go
// Get only district level
result := gc.ReverseGeocodeLevel(-6.2088, 106.8456, geocoder.LevelDistrict)

if result.District != nil {
    fmt.Printf("District: %s\n", result.District.Name)
}
```
