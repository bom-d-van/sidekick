package sidekick

import (
	"flag"
	"log"
	"os"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestSkipCase(t *testing.T) {
	oldArgs := os.Args
	for _, c := range []struct {
		input        []string
		intResult    map[int]bool
		stringResult map[string]bool
	}{
		{
			input:        []string{"-skip", "1", "-skip", "2"},
			intResult:    map[int]bool{0: true, 3: true},
			stringResult: map[string]bool{"0": true, "3": true},
		},
		{
			input:        []string{"-case", "1", "-case", "2"},
			intResult:    map[int]bool{1: true, 2: true},
			stringResult: map[string]bool{"1": true, "2": true},
		},
	} {
		os.Args = append(oldArgs, c.input...)
		skips = []string{}
		cases = []string{}
		flag.Parse()

		intResult := map[int]bool{}
		for i := range []int{0, 1, 2, 3} {
			if SkipCase(i) {
				continue
			}
			intResult[i] = true
		}
		if got, want := intResult, c.intResult; !reflect.DeepEqual(got, want) {
			t.Errorf("intResult = %#v; want %#v", got, want)
		}

		stringResult := map[string]bool{}
		for _, i := range []string{"0", "1", "2", "3"} {
			if SkipCase(i) {
				continue
			}
			stringResult[i] = true
		}
		if got, want := stringResult, c.stringResult; !reflect.DeepEqual(got, want) {
			t.Errorf("stringResult = %#v; want %#v", got, want)
		}
	}
}

func ExampleSkipCase() {
	// skip case or cases: go test -skip 1 -skip 2 --> skip case 1 and 2
	// Or: go test -case 1 -case 2 --> run case 1 and 2
	for i, c := range cases {
		if SkipCase(i) {
			continue
		}
		// you can filter cases not needed when debugging tests
		_ = c
	}
}
