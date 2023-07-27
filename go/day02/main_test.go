package main

import (
	"strings"
	"testing"
)

var test_input = "A Y\nB X\nC Z\n"

func Test_part1(t *testing.T) {
	expected := 15

	ans := part1(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}

func Test_part2(t *testing.T) {
	expected := 12

	ans := part2(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}
