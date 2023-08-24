package main

import (
	"strings"
	"testing"
)

var test_input = "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n"

func Test_part1(t *testing.T) {
	expected := 2

	ans := part1(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}

func Test_part2(t *testing.T) {
	expected := 4

	ans := part2(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}
