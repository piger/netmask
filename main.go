package main

import (
	"errors"
	"fmt"
	"net/netip"
	"os"
)

func printNetmask(prefix netip.Prefix) {
	masked := prefix.Masked()

	for ip := masked.Addr(); masked.Contains(ip); ip = ip.Next() {
		fmt.Println(ip)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <netmask> [address]\n", os.Args[0])
}

func run() error {
	narg := len(os.Args)

	if narg < 2 {
		usage()
		return errors.New("error: not enough arguments")
	}

	prefix, err := netip.ParsePrefix(os.Args[1])
	if err != nil {
		return err
	}

	switch narg {
	case 2:
		printNetmask(prefix)
	case 3:
		ip, err := netip.ParseAddr(os.Args[2])
		if err != nil {
			return err
		}
		if !prefix.Contains(ip) {
			return fmt.Errorf("%s does not contain %s", prefix.Masked(), ip)
		}
	default:
		usage()
		return errors.New("error: too many arguments")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
