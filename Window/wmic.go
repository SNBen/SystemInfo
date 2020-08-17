package window

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"os/exec"
	"regexp"
)

/**
 * 获取电脑CPUId
 */
func GetCpuId() string {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	str := string(out)
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:]
}

//BIOS信息
func GetBiosInfo() string {
	var s []struct {
		Name         string
		SerialNumber string
	}
	err := wmi.Query("SELECT Name , SerialNumber FROM Win32_BIOS WHERE (Name IS NOT NULL)", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return ""
	}
	return s[0].Name
}

//BIOS信息
func GetSerialNumber() string {
	var s []struct {
		Name         string
		SerialNumber string
	}
	err := wmi.Query("SELECT Name ,SerialNumber FROM Win32_BIOS WHERE (SerialNumber IS NOT NULL)", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return ""
	}
	return s[0].SerialNumber
}

//主板信息
func GetMotherboardInfo() string {
	var s []struct {
		Product      string
		SerialNumber string
		Manufacturer string
	}
	err := wmi.Query("SELECT Product,SerialNumber, Manufacturer  FROM Win32_BaseBoard WHERE (Product IS NOT NULL)", &s)
	if err != nil {
		return ""
	}
	if len(s) == 0{
		return ""
	}else {
		return s[0].SerialNumber
	}
}
