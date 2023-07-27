package main

import (
	"strings"
	"testing"
)

var test_input = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000\n"

func Test_part1(t *testing.T) {
	expected := 24000

	ans := part1(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}

func Test_part2(t *testing.T) {
	expected := 45000

	ans := part2(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}
