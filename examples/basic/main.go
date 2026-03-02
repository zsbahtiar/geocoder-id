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

	result := gc.ReverseGeocode(-6.2088, 106.8456)

	if result.Province != nil {
		fmt.Printf("Province: %s (%s)\n", result.Province.Name, result.Province.Code)
	}
	if result.Regency != nil {
		fmt.Printf("Regency:  %s (%s)\n", result.Regency.Name, result.Regency.Code)
	}
	if result.District != nil {
		fmt.Printf("District: %s (%s)\n", result.District.Name, result.District.Code)
	}
	if result.Village != nil {
		fmt.Printf("Village:  %s (%s)\n", result.Village.Name, result.Village.Code)
	}
}
