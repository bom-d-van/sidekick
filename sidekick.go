// Package sidekick implements some Go test helper flags and functions.
package sidekick

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

type stringSlice []string

func (ss stringSlice) String() string      { return strings.Join(ss, ",") }
func (ss *stringSlice) Set(s string) error { *ss = append(*ss, s); return nil }

var (
	// Debug is similar to testing.Verbose. Set to true when using flag
	// -debug.
	Debug bool

	skips stringSlice
	cases stringSlice
)

func init() {
	flag.Var(&skips, "skip", "skip test case")
	flag.Var(&cases, "case", "run test case")
	flag.BoolVar(&Debug, "debug", false, "enter debug mode")

	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// SkipCase checks if a test case should be run based on flags -skip or -case.
// SkipCase compares values in string form, you can pass in value has
// fmt.Stringer or any other interfaces that package fmt has supported.
func SkipCase(c interface{}) bool {
	for _, item := range skips {
		if item == fmt.Sprint(c) {
			return true
		}
	}

	for _, item := range cases {
		if item == fmt.Sprint(c) {
			return false
		}
	}

	if len(cases) > 0 {
		return true
	}

	return false
}

// TODO
// type Skiper interface{}
// type Caser interface{}
