package main

import (
	"strings"
	"testing"
)

var test_input = "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw\n"

func Test_part1(t *testing.T) {
	expected := 157

	ans := part1(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}

func Test_part2(t *testing.T) {
	expected := 70

	ans := part2(strings.TrimRight(test_input, "\n"))

	if ans != expected {
		t.Fatalf("%d != %d", ans, expected)
	}
}
