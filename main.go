package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func printNetmask(netmask string) error {
	addrs, err := addressesFromCIDR(netmask)
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	return nil
}

func netmaskContains(netmask, ip string) error {
	addrs, err := addressesFromCIDR(netmask)
	if err != nil {
		return err
	}

	for _, addr := range addrs {
		if ip == addr {
			return nil
		}
	}

	return fmt.Errorf("%s not in %s", ip, netmask)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <netmask> [address]\n", os.Args[0])
	flag.PrintDefaults()
}

func run() error {
	flag.Usage = usage
	flag.Parse()

	switch len(flag.Args()) {
	case 1:
		return printNetmask(flag.Arg(0))
	case 2:
		if ip := net.ParseIP(flag.Arg(1)); ip == nil {
			return fmt.Errorf("cannot parse IP address %q", flag.Arg(1))
		}
		return netmaskContains(flag.Arg(0), flag.Arg(1))
	default:
		usage()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
