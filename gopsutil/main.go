package main

import (
	"time"
	//"github.com/shirou/gopsutil"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	fmt.Println("###########################################################################################")
	//c,_ := cpu.Info()   // cpu信息

	c, _ := cpu.Times(true) // 目前参数没有意义，获取全部cpu的使用信息
	fmt.Println(c)
	fmt.Println("###########################################################################################")

	n, _ := net.IOCounters(true) // 每个网卡的情况(不是网速，获取 /proc/net/dev文件数据)
	fmt.Println(n)
	fmt.Println("###########################################################################################")

	n1, _ := net.IOCounters(false) // 全部网卡的情况
	fmt.Println(n1)
	fmt.Println("###########################################################################################")

	avg, _ := load.Avg()
	fmt.Println(avg)
	fmt.Println("###########################################################################################")

	misc, _ := load.Misc()
	fmt.Println(misc)
	fmt.Println("###########################################################################################")

	// http://www.blogjava.net/fjzag/articles/317773.html
	cpus, _ := cpu.Times(true) // 目前参数没有意义，获取全部cpu的使用信息

	var stat cpu.TimesStat
	for i := range cpus {
		fmt.Printf("%+v", cpus[i])
		stat.User += cpus[i].User // 从系统启动开始累计到当前时刻，处于用户态的运行时
		stat.System += cpus[i].System
		stat.Idle += cpus[i].Idle
		stat.Nice += cpus[i].Nice
		stat.Iowait += cpus[i].Iowait
		stat.Irq += cpus[i].Irq
		stat.Softirq += cpus[i].Softirq
		stat.Steal += cpus[i].Steal
		stat.Guest += cpus[i].Guest
		stat.GuestNice += cpus[i].GuestNice
		stat.Stolen += cpus[i].Stolen
	}
	fmt.Println("")
	fmt.Printf("%+v", stat)
	idle := stat.Idle / (stat.Idle + stat.System + stat.User)
	fmt.Println("")
	fmt.Printf("%+v", idle)

	fmt.Println("每个cpu的使用百分比")
	pers, _ := cpu.Percent(1*time.Second, true)
	fmt.Println(pers)

}
