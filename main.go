package main

import (
	"flag"
	"fmt"
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
			os.Exit(0)
		}
	}

	os.Exit(1)
	return nil
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
		return netmaskContains(flag.Arg(0), flag.Arg(1))
	default:
		usage()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
