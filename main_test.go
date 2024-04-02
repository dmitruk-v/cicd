package main

import "testing"

func TestSum(t *testing.T) {
	type testcase struct {
		name string
		a    int
		b    int
		sum  int
	}
	table := []testcase{
		{name: "1+2=3", a: 1, b: 2, sum: 3},
		{name: "0+2=2", a: 0, b: 2, sum: 2},
		{name: "1+0=1", a: 1, b: 0, sum: 1},
		{name: "-1+2=1", a: -1, b: 2, sum: 1},
		{name: "1+-2=-1", a: 1, b: -2, sum: -1},
		{name: "-1+-2=-3", a: -1, b: -2, sum: -3},
	}
	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			got := sum(tc.a, tc.b)
			if got != tc.sum {
				t.Fatalf("want: %v, got: %v", tc.sum, got)
			}
		})
	}
}
