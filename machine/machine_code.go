/*
author: superl[N.S.T]
github: https://github.com/super-l/
desc: 获取操作系统的相关硬件基础编码信息
*/
package machine

import (
	"errors"
	"net"
	"runtime"
	"strings"

	"github.com/bdgca-wjp/machine-code/machine/os"
	"github.com/bdgca-wjp/machine-code/machine/types"
)

func GetMachineData() (data types.Information) {
	var osMachine OsMachineInterface
	if runtime.GOOS == "darwin" {
		osMachine = os.MacMachine{}
	} else if runtime.GOOS == "linux" {
		osMachine = os.LinuxMachine{}
	} else if runtime.GOOS == "windows" {
		osMachine = os.WindowsMachine{}
	}
	var machineData = osMachine.GetMachine()
	machineData.LocalMacInfo, _ = GetMACAddress()
	return machineData
}

func GetBoardSerialNumber() (data string, err error) {
	var osMachine OsMachineInterface
	if runtime.GOOS == "darwin" {
		osMachine = os.MacMachine{}
	} else if runtime.GOOS == "linux" {
		osMachine = os.LinuxMachine{}
	} else if runtime.GOOS == "windows" {
		osMachine = os.WindowsMachine{}
	}
	return osMachine.GetBoardSerialNumber()
}

func GetPlatformUUID() (data string, err error) {
	var osMachine OsMachineInterface
	if runtime.GOOS == "darwin" {
		osMachine = os.MacMachine{}
	} else if runtime.GOOS == "linux" {
		osMachine = os.LinuxMachine{}
	} else if runtime.GOOS == "windows" {
		osMachine = os.WindowsMachine{}
	}
	return osMachine.GetPlatformUUID()
}

func GetCpuSerialNumber() (data string, err error) {
	var osMachine OsMachineInterface
	if runtime.GOOS == "darwin" {
		osMachine = os.MacMachine{}
	} else if runtime.GOOS == "linux" {
		osMachine = os.LinuxMachine{}
	} else if runtime.GOOS == "windows" {
		osMachine = os.WindowsMachine{}
	}
	return osMachine.GetCpuSerialNumber()
}

func GetMACAddress() ([]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var wlanmac string
	var macs []string
	var runmacs []string

	for i := 0; i < len(netInterfaces); i++ {
		flags := netInterfaces[i].Flags.String()
		//fmt.Printf("全部属性:%+v\n", netInterfaces[i])
		//fmt.Println("Flags:", flags, " Name:", netInterfaces[i].Name, " Mac:", netInterfaces[i].HardwareAddr.String())

		if strings.Contains(flags, "broadcast") &&
			!strings.Contains(flags, "loopback") {

			name := netInterfaces[i].Name
			if strings.Contains(name, "WLAN") {
				if netInterfaces[i].HardwareAddr.String() != "" {
					wlanmac = netInterfaces[i].HardwareAddr.String()
				}
			} else if !strings.Contains(name, "VMware") && //排除VM虚拟机网络
				!strings.Contains(name, "(WSL)") && //排除WSL虚拟机网络
				!strings.Contains(name, "蓝牙") && //排除蓝牙网络
				!strings.Contains(strings.ToUpper(name), "BLUETOOTH") {
				if netInterfaces[i].HardwareAddr.String() != "" {
					if strings.Contains(flags, "running") {
						runmacs = append(runmacs, netInterfaces[i].HardwareAddr.String())
					} else {
						macs = append(macs, netInterfaces[i].HardwareAddr.String())
					}
				}
			}
		}
	}
	if wlanmac != "" {
		arr := []string{wlanmac}
		runmacs = append(arr, runmacs...)
	}
	macs = append(runmacs, macs...)
	if len(macs) == 0 {
		return nil, errors.New("unable to get the correct MAC address")
	}

	return macs, errors.New("unable to get the correct MAC address")
}

func GetLocalIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip, nil
}

func GetIpAddrAll() ([]string, error) {
	var ipList []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ipList, err
	}
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() {
			if ipNet.IP.To4() != nil {
				ipList = append(ipList, ipNet.IP.To4().String())
			}
		}
	}
	return ipList, nil
}
