# API Reference

Geocoder ID provides two interfaces:

## [CLI](/api/cli)

Command-line interface for quick geocoding tasks.

```bash
geocoder geocode --coords="-6.2088 106.8456"
```

## [Go Library](/api/go)

Importable Go package for integration into your applications.

```go
import geocoder "github.com/zsbahtiar/geocoder-id"

gc, _ := geocoder.NewDefault()
result := gc.ReverseGeocode(-6.2088, 106.8456)
```
