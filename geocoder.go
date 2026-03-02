package geocoder

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/duckdb/duckdb-go/v2"
)

const defaultDBName = "geocoder.duckdb"

type Geocoder struct {
	db *sql.DB
}

func New(dbPath string) (*Geocoder, error) {
	db, err := sql.Open("duckdb", dbPath+"?access_mode=read_only")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if _, err := db.Exec("LOAD spatial"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to load spatial extension: %w", err)
	}

	return &Geocoder{db: db}, nil
}

func NewDefault() (*Geocoder, error) {
	dbPath, err := findDatabase()
	if err != nil {
		dbPath, err = autoDownloadDatabase()
		if err != nil {
			return nil, err
		}
	}
	return New(dbPath)
}

func findDatabase() (string, error) {
	candidates := []string{}

	if envPath := os.Getenv("GEOCODER_DB_PATH"); envPath != "" {
		candidates = append(candidates, envPath)
	}

	candidates = append(candidates, defaultDBName)
	candidates = append(candidates, filepath.Join("data", defaultDBName))

	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		candidates = append(candidates, filepath.Join(exeDir, defaultDBName))
		candidates = append(candidates, filepath.Join(exeDir, "data", defaultDBName))
	}

	if home, err := os.UserHomeDir(); err == nil {
		candidates = append(candidates, filepath.Join(home, ".geocoder-id", defaultDBName))
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("database not found")
}

func autoDownloadDatabase() (string, error) {
	dbPath, err := getDefaultDBPath()
	if err != nil {
		return "", err
	}

	if err := downloadDatabase(dbPath); err != nil {
		return "", err
	}

	return dbPath, nil
}

func (g *Geocoder) Close() error {
	return g.db.Close()
}

func (g *Geocoder) ReverseGeocode(lat, lon float64) Result {
	return g.ReverseGeocodeLevel(lat, lon, LevelVillage)
}

func (g *Geocoder) ReverseGeocodeLevel(lat, lon float64, level Level) Result {
	return g.queryLevel(lat, lon, string(level))
}

func (g *Geocoder) queryLevel(lat, lon float64, level string) Result {
	var result Result

	query := `
		SELECT h.province_code, h.province_name,
			   h.regency_code, h.regency_name,
			   h.district_code, h.district_name,
			   h.village_code, h.village_name
		FROM locations l
		JOIN hierarchy h ON l.code = h.code
		WHERE ST_Contains(l.geom, ST_Point(?, ?))
		  AND l.level = ?
		LIMIT 1
	`

	row := g.db.QueryRow(query, lon, lat, level)

	var provCode, provName sql.NullString
	var regCode, regName sql.NullString
	var distCode, distName sql.NullString
	var vilCode, vilName sql.NullString

	err := row.Scan(&provCode, &provName, &regCode, &regName, &distCode, &distName, &vilCode, &vilName)
	if err != nil {
		return result
	}

	if provCode.Valid {
		result.Province = &Location{Code: provCode.String, Name: provName.String}
	}
	if regCode.Valid {
		result.Regency = &Location{Code: regCode.String, Name: regName.String}
	}
	if distCode.Valid {
		result.District = &Location{Code: distCode.String, Name: distName.String}
	}
	if vilCode.Valid {
		result.Village = &Location{Code: vilCode.String, Name: vilName.String}
	}

	return result
}

func ParseCoords(s string) []Coord {
	var coords []Coord
	parts := strings.Split(s, ";")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		fields := strings.Fields(part)
		if len(fields) != 2 {
			continue
		}

		lat, err1 := strconv.ParseFloat(fields[0], 64)
		lon, err2 := strconv.ParseFloat(fields[1], 64)
		if err1 != nil || err2 != nil {
			continue
		}

		coords = append(coords, Coord{Lat: lat, Lon: lon})
	}

	return coords
}
