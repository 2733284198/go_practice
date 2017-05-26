package main

import (
	//"github.com/shirou/gopsutil"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/load"
	"fmt"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	fmt.Println("###########################################################################################")
	//c,_ := cpu.Info()   // cpu信息

	c,_:=cpu.Times(true)  // 目前参数没有意义，获取全部cpu的使用信息
	fmt.Println(c)
	fmt.Println("###########################################################################################")

	n,_ := net.IOCounters(true)	// 每个网卡的情况
	fmt.Println(n)
	fmt.Println("###########################################################################################")

	n1,_ := net.IOCounters(false)	// 全部网卡的情况
	fmt.Println(n1)
	fmt.Println("###########################################################################################")

	avg,_ := load.Avg()
	fmt.Println(avg)
	fmt.Println("###########################################################################################")

	misc,_ := load.Misc()
	fmt.Println(misc)
	fmt.Println("###########################################################################################")
}
