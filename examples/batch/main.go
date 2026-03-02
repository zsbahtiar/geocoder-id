package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	geocoder "github.com/zsbahtiar/geocoder-id"
)

func main() {
	gc, err := geocoder.NewDefault()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.Close()

	coordinates := []geocoder.Coord{
		{Lat: -6.2088, Lon: 106.8456}, // Jakarta
		{Lat: -6.9175, Lon: 107.6191}, // Bandung
		{Lat: -7.2575, Lon: 112.7521}, // Surabaya
		{Lat: -8.6705, Lon: 115.2126}, // Bali
		{Lat: -2.0498, Lon: 110.1879}, // Kalimantan Barat
	}

	var results []geocoder.Result
	for _, coord := range coordinates {
		result := gc.ReverseGeocode(coord.Lat, coord.Lon)
		results = append(results, result)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(results)

	fmt.Fprintf(os.Stderr, "\nProcessed %d coordinates\n", len(coordinates))
}
