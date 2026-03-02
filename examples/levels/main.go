package main

import (
	"fmt"
	"log"

	geocoder "github.com/zsbahtiar/geocoder-id"
)

func main() {
	gc, err := geocoder.NewDefault()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.Close()

	lat, lon := -6.2088, 106.8456

	fmt.Println("=== Query at different levels ===\n")

	levels := []geocoder.Level{
		geocoder.LevelProvince,
		geocoder.LevelRegency,
		geocoder.LevelDistrict,
		geocoder.LevelVillage,
	}

	for _, level := range levels {
		result := gc.ReverseGeocodeLevel(lat, lon, level)
		fmt.Printf("Level: %s\n", level)
		if result.Province != nil {
			fmt.Printf("  Result: %s", result.Province.Name)
			if result.Regency != nil {
				fmt.Printf(" > %s", result.Regency.Name)
			}
			if result.District != nil {
				fmt.Printf(" > %s", result.District.Name)
			}
			if result.Village != nil {
				fmt.Printf(" > %s", result.Village.Name)
			}
			fmt.Println()
		} else {
			fmt.Println("  Result: not found")
		}
		fmt.Println()
	}
}
