package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	geocoder "github.com/zsbahtiar/geocoder-id"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "geocoder",
		Short:   "Offline reverse geocoding for Indonesia",
		Long:    `Geocoder ID | Convert coordinates (lat/lon) to Indonesia's administrative addresses up to the village/sub-district level`,
		Version: geocoder.Version,
	}

	rootCmd.AddCommand(geocodeCmd())

	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func geocodeCmd() *cobra.Command {
	var coords, dbPath, output, level string
	var debug bool

	cmd := &cobra.Command{
		Use:   "geocode",
		Short: "Reverse geocoding coordinates to addresses",
		Example: `  geocoder geocode --coords="-6.2088 106.8456"
  geocoder geocode --coords="-6.2088 106.8456;-2.04963 110.18774"
  geocoder geocode --coords="-6.2088 106.8456" --output=json`,
		Run: func(cmd *cobra.Command, args []string) {
			log := func(format string, a ...interface{}) {
				if debug {
					ts := time.Now().Format("2006-01-02 15:04:05.000")
					fmt.Fprintf(os.Stderr, "%s [DEBUG] "+format+"\n", append([]interface{}{ts}, a...)...)
				}
			}

			log("Starting geocode process")
			log("Database: %s", dbPath)
			log("Level: %s", level)
			log("Output: %s", output)

			log("Parsing coordinates: %s", coords)
			coordList := geocoder.ParseCoords(coords)
			if len(coordList) == 0 {
				fmt.Fprintln(os.Stderr, "Invalid coords format. Use: \"lat lon\" or \"lat lon;lat lon;...\"")
				os.Exit(1)
			}
			log("Parsed %d coordinate(s)", len(coordList))

			log("Opening database...")
			var gc *geocoder.Geocoder
			var err error
			if dbPath != "" {
				gc, err = geocoder.New(dbPath)
			} else {
				gc, err = geocoder.NewDefault()
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			defer gc.Close()
			log("Database opened successfully")

			var results []geocoder.Result
			for i, c := range coordList {
				log("Processing coordinate %d: lat=%f, lon=%f", i+1, c.Lat, c.Lon)
				log("Querying level: %s", level)

				result := gc.ReverseGeocodeLevel(c.Lat, c.Lon, geocoder.Level(level))

				if result.Province == nil {
					log("Coordinate %d: no result found", i+1)
				} else {
					log("Coordinate %d: found %s", i+1, result.Province.Name)
				}
				results = append(results, result)
			}

			var nonEmptyResults []geocoder.Result
			for _, r := range results {
				if r.Province != nil {
					nonEmptyResults = append(nonEmptyResults, r)
				}
			}

			log("Geocoding complete, %d result(s) found out of %d coordinate(s)", len(nonEmptyResults), len(results))

			if len(nonEmptyResults) == 0 {
				log("No results to output")
			} else if output == "json" {
				printJSON(nonEmptyResults)
			} else {
				printTable(nonEmptyResults)
			}

			log("Done")
		},
	}

	cmd.Flags().StringVar(&coords, "coords", "", "Coordinates: \"lat lon\" or \"lat lon;lat lon;...\"")
	cmd.Flags().StringVar(&dbPath, "db", "", "Database path (auto-detect if not specified)")
	cmd.Flags().StringVar(&output, "output", "table", "Output format: table, json")
	cmd.Flags().StringVar(&level, "level", "village", "Level: province, regency, district, village")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug logging")
	cmd.MarkFlagRequired("coords")

	return cmd
}

func printJSON(results []geocoder.Result) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(results)
}

func printTable(results []geocoder.Result) {
	headers := []string{"province_code", "province_name", "regency_code", "regency_name", "district_code", "district_name", "village_code", "village_name"}
	widths := []int{13, 13, 12, 12, 13, 13, 12, 12}

	var rows [][]string
	for _, result := range results {
		values := make([]string, 8)
		if result.Province != nil {
			values[0] = result.Province.Code
			values[1] = result.Province.Name
		}
		if result.Regency != nil {
			values[2] = result.Regency.Code
			values[3] = result.Regency.Name
		}
		if result.District != nil {
			values[4] = result.District.Code
			values[5] = result.District.Name
		}
		if result.Village != nil {
			values[6] = result.Village.Code
			values[7] = result.Village.Name
		}
		rows = append(rows, values)
	}

	for i, h := range headers {
		if len(h) > widths[i] {
			widths[i] = len(h)
		}
	}
	for _, row := range rows {
		for i, v := range row {
			if len(v) > widths[i] {
				widths[i] = len(v)
			}
		}
	}

	var hdr, sep strings.Builder
	hdr.WriteString("|")
	sep.WriteString("|")

	for i := range headers {
		hdr.WriteString(fmt.Sprintf(" %-*s |", widths[i], headers[i]))
		sep.WriteString(strings.Repeat("-", widths[i]+2) + "|")
	}

	fmt.Println(hdr.String())
	fmt.Println(sep.String())

	for _, row := range rows {
		var line strings.Builder
		line.WriteString("|")
		for i := range row {
			line.WriteString(fmt.Sprintf(" %-*s |", widths[i], row[i]))
		}
		fmt.Println(line.String())
	}
}
