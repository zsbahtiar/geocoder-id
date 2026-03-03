# Go Library Reference

```go
import geocoder "github.com/zsbahtiar/geocoder-id"
```

## Types

### Geocoder

Main geocoder instance.

```go
type Geocoder struct {
    // contains filtered or unexported fields
}
```

### Result

Geocoding result containing all administrative levels.

```go
type Result struct {
    Province *Location `json:"province,omitempty"`
    Regency  *Location `json:"regency,omitempty"`
    District *Location `json:"district,omitempty"`
    Village  *Location `json:"village,omitempty"`
}
```

### Location

Administrative location with code and name.

```go
type Location struct {
    Code string `json:"code"`
    Name string `json:"name"`
}
```

### Coord

Coordinate pair.

```go
type Coord struct {
    Lat float64
    Lon float64
}
```

### Level

Administrative level constant.

```go
type Level string

const (
    LevelProvince Level = "province"
    LevelRegency  Level = "regency"
    LevelDistrict Level = "district"
    LevelVillage  Level = "village"
)
```

## Functions

### NewDefault

Creates a new Geocoder with auto-detected database path. Downloads database if not found.

```go
func NewDefault() (*Geocoder, error)
```

**Example:**
```go
gc, err := geocoder.NewDefault()
if err != nil {
    log.Fatal(err)
}
defer gc.Close()
```

### New

Creates a new Geocoder with explicit database path.

```go
func New(dbPath string) (*Geocoder, error)
```

**Example:**
```go
gc, err := geocoder.New("/path/to/geocoder.duckdb")
if err != nil {
    log.Fatal(err)
}
defer gc.Close()
```

### ParseCoords

Parses coordinate string into slice of Coord.

```go
func ParseCoords(s string) []Coord
```

**Example:**
```go
coords := geocoder.ParseCoords("-6.2088 106.8456;-7.7956 110.3695")
// Returns: []Coord{{Lat: -6.2088, Lon: 106.8456}, {Lat: -7.7956, Lon: 110.3695}}
```

## Methods

### (*Geocoder) ReverseGeocode

Reverse geocode coordinates to village level.

```go
func (g *Geocoder) ReverseGeocode(lat, lon float64) Result
```

**Example:**
```go
result := gc.ReverseGeocode(-6.2088, 106.8456)
if result.Village != nil {
    fmt.Println(result.Village.Name) // "Pasar Manggis"
}
```

### (*Geocoder) ReverseGeocodeLevel

Reverse geocode coordinates to specific administrative level.

```go
func (g *Geocoder) ReverseGeocodeLevel(lat, lon float64, level Level) Result
```

**Example:**
```go
result := gc.ReverseGeocodeLevel(-6.2088, 106.8456, geocoder.LevelDistrict)
if result.District != nil {
    fmt.Println(result.District.Name) // "Setiabudi"
}
```

### (*Geocoder) Close

Closes the database connection. Always call this when done.

```go
func (g *Geocoder) Close() error
```

**Example:**
```go
gc, _ := geocoder.NewDefault()
defer gc.Close()
```
