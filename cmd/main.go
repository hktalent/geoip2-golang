package main

import (
	"fmt"
	"github.com/hktalent/geoip2-golang"
)

func main() {
	// If you are using strings that may be invalid, check that ip is not nil
	if x, err := geoip2.GetIpsCity("81.2.69.142", "110.242.68.66", "39.156.66.10"); nil == err {
		for record := range x {
			fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
			if len(record.Subdivisions) > 0 {
				fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
			}
			fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
			fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
			fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
			fmt.Printf("Coordinates: %v, %v\n====================\n", record.Location.Latitude, record.Location.Longitude)
		}
	}
	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
	geoip2.Close()
}
