package window

import (
	"fmt"
	"os/exec"
	"strings"
)

func PsInfo() map[string]string {
	//System information for \\BEN-PC:
	//Uptime:                    0 days 2 hours 8 minutes 11 seconds
	//Kernel version:            Windows 7 Ultimate, Multiprocessor Free
	//Product type:              Professional
	//Product version:           6.1
	//Service pack:              0
	//Kernel build number:       7601
	//Registered organization:   Microsoft
	//Registered owner:          Microsoft
	//IE version:                9.0000
	//System root:               C:\Windows
	//Processors:                8
	//Processor speed:           2.8 GHz
	//Processor type:            Intel(R) Core(TM) i7-7700HQ CPU @
	//Physical memory:           34 MB
	//Video driver:              Intel(R) HD Graphics 630
	cmd := exec.Command("psinfo.exe", "-nobanner")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	str := string(out)

	s := strings.Split(str, "\r\n")
	var Result = map[string]string{}

	for _, value := range s {
		if len(value) == 0 {
			continue
		}
		if strings.Contains(value, "System information for") {
			continue
		}
		kv := strings.Split(value, ":")
		//fmt.Println(kv[0], kv[1])
		Result[kv[0]] = strings.Trim(kv[1], " ")
	}
	return Result
}

func SystemInfo() string {
	//System information for \\BEN-PC:
	//Uptime:                    0 days 2 hours 8 minutes 11 seconds
	//Kernel version:            Windows 7 Ultimate, Multiprocessor Free
	//Product type:              Professional
	//Product version:           6.1
	//Service pack:              0
	//Kernel build number:       7601
	//Registered organization:   Microsoft
	//Registered owner:          Microsoft
	//IE version:                9.0000
	//System root:               C:\Windows
	//Processors:                8
	//Processor speed:           2.8 GHz
	//Processor type:            Intel(R) Core(TM) i7-7700HQ CPU @
	//Physical memory:           34 MB
	//Video driver:              Intel(R) HD Graphics 630
	cmd := exec.Command("SystemInfo.exe")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	str := string(out)

	return str
}
