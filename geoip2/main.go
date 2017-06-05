package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("60.191.59.162")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	// City Info
	fmtStr := "zh-CN"
	// 城市名字
	fmt.Printf("City Name: %v\n", record.City.Names[fmtStr])
	// 城市编号
	fmt.Printf("City GeoNameID : %d\n",record.City.GeoNameID)
	// 大陆（洲）信息
	fmt.Println("Continent Code : ",record.Continent.Code)
	fmt.Println("Continent GeoNameID : ",record.Continent.GeoNameID)
	fmt.Println("Continent Names : ",record.Continent.Names[fmtStr])

	// 所属国家信息
	fmt.Printf("Country Name: %v\n", record.Country.Names[fmtStr])
	fmt.Printf("Country ISO Code: %v\n", record.Country.IsoCode)
	fmt.Println("Country GeoNameID ：",record.Country.GeoNameID)


	// Postal
	//fmt.Println("Postal Code : ", record.Postal.Code)

	// RegisteredCountry ??
	//fmt.Println("RepresentedCountry : ",record.RegisteredCountry.Names[fmtStr])
	//fmt.Println("RepresentedCountry : ",record.RepresentedCountry.Names[fmtStr])

	// 省信息
	fmt.Println("Subdivisions Name ：",record.Subdivisions[0].Names[fmtStr])
	fmt.Println("Subdivisions GeoNameID ：",record.Subdivisions[0].GeoNameID)
	fmt.Println("Subdivisions IsoCode ：",record.Subdivisions[0].IsoCode)


	fmt.Println("****************************************************************")
	// ISP Info

	db2 , err := geoip2.Open("GeoLite2-ASN.mmdb")
	isp , err := db2.ISP(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(isp)


}