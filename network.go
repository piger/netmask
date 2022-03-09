package main

import (
	"net"
)

// Original code by Russ Cox: https://swtch.com/~rsc/
// https://groups.google.com/g/golang-nuts/c/zlcYA4qk-94?pli=1

func addressesFromCIDR(cidr string) (chan string, error) {
	out := make(chan string)

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return out, err
	}

	go func() {
		for addr := ip.Mask(ipnet.Mask); ipnet.Contains(addr); incrementIP(addr) {
			out <- addr.String()
		}
		close(out)
	}()

	return out, nil
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
