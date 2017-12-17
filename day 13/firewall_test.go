package main

import (
	"testing"
	"fmt"
)

func TestSeverity(t *testing.T) {

	cases := []struct{ depth, length int }{{0, 3}, {1,2},{4,4},{6,4}}
	severity := 0
	for _, f := range cases {
		severity += firewall{f.depth, f.length}.Severity()
	}
	fmt.Println(severity)
}
