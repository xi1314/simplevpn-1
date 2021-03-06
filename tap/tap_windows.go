// +build windows
package tap

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var (
	IfceNameNotFound  = errors.New("Failed to find the name of interface.")
	TapDeviceNotFound = errors.New("Failed to find the tap device in registry.")
	// Device Control Codes
	tap_win_ioctl_get_mac             = tap_control_code(1, 0)
	tap_win_ioctl_get_version         = tap_control_code(2, 0)
	tap_win_ioctl_get_mtu             = tap_control_code(3, 0)
	tap_win_ioctl_get_info            = tap_control_code(4, 0)
	tap_ioctl_config_point_to_point   = tap_control_code(5, 0)
	tap_ioctl_set_media_status        = tap_control_code(6, 0)
	tap_win_ioctl_config_dhcp_masq    = tap_control_code(7, 0)
	tap_win_ioctl_get_log_line        = tap_control_code(8, 0)
	tap_win_ioctl_config_dhcp_set_opt = tap_control_code(9, 0)
	tap_ioctl_config_tun              = tap_control_code(10, 0)
	// w32 api
	file_device_unknown = uint32(0x00000022)
)

func ctl_code(device_type, function, method, access uint32) uint32 {
	return (device_type << 16) | (access << 14) | (function << 2) | method
}

func tap_control_code(request, method uint32) uint32 {
	return ctl_code(file_device_unknown, request, method, 0)
}

// GetDeviceId finds out a TAP device from registry, it requires privileged right.
func getdeviceid() (string, error) {
	// TAP driver key location
	regkey := `SYSTEM\CurrentControlSet\Control\Class\{4D36E972-E325-11CE-BFC1-08002BE10318}`
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey, registry.ALL_ACCESS)
	if err != nil {
		return "", err
	}
	defer k.Close()
	// read all subkeys
	keys, err := k.ReadSubKeyNames(-1)
	if err != nil {
		return "", err
	}
	// find the one with ComponentId == "tap0901"
	for _, v := range keys {
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey+"\\"+v, registry.ALL_ACCESS)
		if err != nil {
			continue
		}
		val, _, err := key.GetStringValue("ComponentId")
		if err != nil {
			goto next
		}
		if val == "tap0901" {
			val, _, err = key.GetStringValue("NetCfgInstanceId")
			if err != nil {
				goto next
			}
			key.Close()
			return val, nil
		}
	next:
		key.Close()
	}
	return "", TapDeviceNotFound
}

// NewTAP find and open a TAP device.
func newTAP() (ifce *Interface, err error) {
	deviceid, err := getdeviceid()
	if err != nil {
		return nil, err
	}
	path := "\\\\.\\Global\\" + deviceid + ".tap"
	pathp, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}
	// type Handle uintptr
	file, err := syscall.CreateFile(pathp, syscall.GENERIC_READ|syscall.GENERIC_WRITE, uint32(syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE), nil, syscall.OPEN_EXISTING, syscall.FILE_ATTRIBUTE_SYSTEM, 0)
	// if err hanppens, close the interface.
	defer func() {
		if err != nil {
			syscall.Close(file)
		}
		if err := recover(); err != nil {
			syscall.Close(file)
		}
	}()
	if err != nil {
		return nil, err
	}
	var bytesReturned uint32
	// find the mac address of tap device.
	mac := make([]byte, 6)
	err = syscall.DeviceIoControl(file, tap_win_ioctl_get_mac, &mac[0], uint32(len(mac)), &mac[0], uint32(len(mac)), &bytesReturned, nil)
	if err != nil {
		return nil, err
	}
	//TUN
	//code2 := []byte{0x0a, 0x03, 0x00, 0x01, 0x0a, 0x03, 0x00, 0x00, 0xff, 0xff, 0xff, 0x00}
	//err = syscall.DeviceIoControl(file, tap_ioctl_config_tun, &code2[0], uint32(12), &rdbbuf[0], uint32(len(rdbbuf)), &bytesReturned, nil)
	//if err != nil {
	//	log.Fatalln("code2 err:", err)
	//}
	fd := os.NewFile(uintptr(file), path)
	ifce = &Interface{tap: true, file: fd}
	copy(ifce.mac[:6], mac[:6])

	// find the name of tap interface(to set the ip)
	hwaddr_equal := func(a net.HardwareAddr, b []byte) bool {
		for i := 0; i < 6; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	ifces, err := net.Interfaces()
	if err != nil {
		return
	}
	// bring up device.
	rdbbuf := make([]byte, syscall.MAXIMUM_REPARSE_DATA_BUFFER_SIZE)
	code := []byte{0x01, 0x00, 0x00, 0x00}
	err = syscall.DeviceIoControl(file, tap_ioctl_set_media_status, &code[0], uint32(4), &rdbbuf[0], uint32(len(rdbbuf)), &bytesReturned, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range ifces {
		if hwaddr_equal(v.HardwareAddr[:6], mac[:6]) {
			ifce.name = v.Name
			//ifce.ignoreDefaultRoutes()
			return
		}
	}
	err = IfceNameNotFound
	return

}

func (ifce *Interface) ignoreDefaultRoutes() error {
	sargs := "interface ip set interface interface=REPLACE_ME ignoredefaultroutes=enabled"
	args := strings.Split(sargs, " ")
	args[4] = fmt.Sprintf("interface=\"%s\"", ifce.name)
	cmd := exec.Command("netsh", args...)
	return cmd.Run()
}

func (ifce *Interface) setIP(ip_mask *net.IPNet) (err error) {
	sargs := fmt.Sprintf("interface ip set address name=REPLACE_ME source=static addr=REPLACE_ME mask=REPLACE_ME gateway=none")
	args := strings.Split(sargs, " ")
	args[4] = fmt.Sprintf("name=\"%s\"", ifce.Name())
	args[6] = fmt.Sprintf("addr=%s", ip_mask.IP)
	args[7] = fmt.Sprintf("mask=%d.%d.%d.%d", ip_mask.Mask[0], ip_mask.Mask[1], ip_mask.Mask[2], ip_mask.Mask[3])
	cmd := exec.Command("netsh", args...)
	err = cmd.Run()
	return
}

func addRoute(ip net.IP, ip_mask *net.IPNet, ifce string) (err error) {
	sargs := fmt.Sprintf("interface ip add route %s REPLACE_ME %s", ip_mask, ip)
	args := strings.Split(sargs, " ")
	args[5] = fmt.Sprintf("\"%s\"", ifce)
	cmd := exec.Command("netsh", args...)
	err = cmd.Run()
	return
}

func (ifce *Interface) addRoute(ip net.IP, ip_mask *net.IPNet) (err error) {
	return addRoute(ip, ip_mask, ifce.name)
}

func delRoute(ip net.IP, ip_mask *net.IPNet, ifce string) (err error) {
	sargs := fmt.Sprintf("interface ip del route %s REPLACE_ME %s", ip_mask, ip)
	args := strings.Split(sargs, " ")
	args[5] = fmt.Sprintf("\"%s\"", ifce)
	cmd := exec.Command("netsh", args...)
	err = cmd.Run()
	return
}

func (ifce *Interface) delRoute(ip net.IP, ip_mask *net.IPNet) (err error) {
	return delRoute(ip, ip_mask, ifce.name)
}
