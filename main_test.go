package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestNetmask(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"netmask": main,
	})
}
