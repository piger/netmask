package main

import (
	"net"
)

// Code by Russ Cox: https://swtch.com/~rsc/
// https://groups.google.com/g/golang-nuts/c/zlcYA4qk-94?pli=1

func addressesFromCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for addr := ip.Mask(ipnet.Mask); ipnet.Contains(addr); incrementIP(addr) {
		ips = append(ips, addr.String())
	}

	return ips, nil
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
