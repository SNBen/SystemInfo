//package SystemInfo
package main

import (
	"SystemInfo/Window"
	"fmt"
	"runtime"
)

func main() {
	sysInfo := window.GetCpuInfo()
	TextWinDev()
	fmt.Println(sysInfo)
}

//func SystemInfo() map[string]string {
//    return window.PsInfo()
//}

func TextWinDev() {
	//fmt.Printf("%-8s : %s\n","开机时长", GetStartTime())
	fmt.Printf("%-8s : %s\n", "当前用户", window.GetUserName())
	fmt.Printf("%-8s : %s\n", "当前系统", runtime.GOOS)
	fmt.Printf("%-8s : %s\n", "系统版本", window.GetSystemVersion())
	fmt.Printf("%-12s : %s\n", "Bios Ver", window.GetBiosInfo())
	fmt.Printf("%-20s : %s\n", "Motherboard(S/N)", window.GetMotherboardInfo())
	fmt.Printf("%-20s : %s\n", "SerialNumber(Type)", window.GetSerialNumber())
	fmt.Printf("%-12s : %s\n", "CPU", window.GetCpuInfo())
	fmt.Printf("%-12s : %s MB\n", "Memory", window.GetMemory())
	fmt.Printf("%-12s : \n %v\n", "Disk", window.GetDiskInfo())
	fmt.Printf("%-12s : \n %v\n", "Interfaces", window.GetIntfs())
}
