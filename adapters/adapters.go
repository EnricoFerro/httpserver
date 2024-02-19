package adapters

import (
	"net"
	"os"
	"syscall"
	"unsafe"
)

type AdapterWithIPNets struct {
	Name        string
	Description string
	IPNets      []net.IPNet
}

func Adapters() ([]AdapterWithIPNets, error) {
	var awins []AdapterWithIPNets
	ai, err := getAdapterList()
	if err != nil {
		return nil, err
	}
	for ; ai != nil; ai = ai.Next {
		name := bytePtrToString(&ai.AdapterName[0])
		description := bytePtrToString(&ai.Description[0])
		awin := AdapterWithIPNets{Name: name, Description: description}
		iai := &ai.IpAddressList
		for ; iai != nil; iai = iai.Next {
			ip := net.ParseIP(bytePtrToString(&iai.IpAddress.String[0]))
			mask := parseIPv4Mask(bytePtrToString(&iai.IpMask.String[0]))
			awin.IPNets = append(awin.IPNets, net.IPNet{IP: ip, Mask: mask})
		}
		awins = append(awins, awin)
	}
	return awins, nil
}

func parseIPv4Mask(ipStr string) net.IPMask {
	ip := net.ParseIP(ipStr).To4()
	return net.IPv4Mask(ip[0], ip[1], ip[2], ip[3])
}

// https://github.com/golang/go/blob/go1.4.1/src/net/interface_windows.go#L13-L20
func bytePtrToString(p *uint8) string {
	a := (*[10000]uint8)(unsafe.Pointer(p))
	i := 0
	for a[i] != 0 {
		i++
	}
	return string(a[:i])
}

// copied from https://github.com/golang/go/blob/go1.4.1/src/net/interface_windows.go#L22-L39
func getAdapterList() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	// TODO(mikio): GetAdaptersInfo returns IP_ADAPTER_INFO that
	// contains IPv4 address list only. We should use another API
	// for fetching IPv6 stuff from the kernel.
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}
