package main

import(
	"github.com/kellydunn/golang-geo"
	"fmt"
)

func main() {
     // Make a few points
     p := geo.NewPoint(34.5, 121.43333)		// 上海
     p2 := geo.NewPoint(30.26667, 120.2)	// 杭州
	 p3 := geo.NewPoint(23.16667,113.23333)	// 广州
	 p4 := geo.NewPoint(25.03,121.30)		// 台北
	 p5 := geo.NewPoint(22.0000,113.5)      // 澳门
     // find the great circle distance between them
     dist := p.GreatCircleDistance(p2)
	 fmt.Printf("上海 --> 杭州 : %+v\n", dist)

	 dist = p.GreatCircleDistance(p3)
	 fmt.Printf("上海 --> 广州 : %+v\n", dist)

	 dist = p.GreatCircleDistance(p5)
	 fmt.Printf("上海 --> 澳门 : %+v\n", dist)

	 dist = p.GreatCircleDistance(p4)
	 fmt.Printf("上海 --> 台北 : %+v\n", dist)

	 dist = p2.GreatCircleDistance(p4)
	 fmt.Printf("杭州 --> 台北 : %+v\n", dist)

	 dist = p3.GreatCircleDistance(p4)
	 fmt.Printf("广州 --> 台北 : %+v\n", dist)
}