package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var input string

func init() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimRight(string(content), "\n")
}

func main() {
	ans1 := part1(input)
	fmt.Println("Part 1: ", ans1)
	ans2 := part2(input)
	fmt.Println("Part 2: ", ans2)
}

func part1(input string) int {
	result := 0

	for _, bag := range strings.Split(input, "\n") {
		set := map[string]bool{}
		duplicates := map[string]bool{}

		for i, r := range strings.Split(bag, "") {
			if i < len(bag)/2 {
				set[r] = true
			} else {
				if set[r] {
					if !duplicates[r] {
						duplicates[r] = true
					}
				}
			}
		}
		for x := range duplicates {
			ascii := int(rune(x[0]))
			if ascii >= 97 {
				result += ascii - 96
			} else {
				result += ascii - 38
			}
		}
	}
	return result
}

func part2(input string) int {
	result := 0

	threes := []string{}
	sets := [3]map[string]bool{{}, {}, {}}

	for _, line := range strings.Split(input, "\n") {
		threes = append(threes, line)
		if len(threes) == 3 {
			for i, bag := range threes {
				for _, c := range strings.Split(bag, "") {
					sets[i][c] = true
				}
			}

			for k := range sets[0] {
				if sets[1][k] && sets[2][k] {
					ascii := int(rune(k[0]))
					if ascii >= 97 {
						result += ascii - 96
					} else {
						result += ascii - 38
					}
				}
			}
			threes = []string{}
			sets = [3]map[string]bool{{}, {}, {}}
		}
	}
	return result
}
