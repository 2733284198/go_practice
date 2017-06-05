package main

import (
	"fmt"
	"github.com/ip2location/ip2location-go"
)

func main() {
	ip2location.Open("./IP-COUNTRY-REGION-CITY-ISP-SAMPLE.BIN")
	ip := "60.191.59.162"
	
	results := ip2location.Get_all(ip)
	
	fmt.Printf("country_short: %s\n", results.Country_short)
	fmt.Printf("country_long: %s\n", results.Country_long)
	fmt.Printf("region: %s\n", results.Region)
	fmt.Printf("city: %s\n", results.City)
	fmt.Printf("isp: %s\n", results.Isp)
	fmt.Printf("api version: %s\n", ip2location.Api_version())
	
	ip2location.Close()
}