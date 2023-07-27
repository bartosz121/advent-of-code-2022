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
	score := 0

	for _, game := range strings.Split(input, "\n") {
		current := 0
		enemy, me := func() (string, string) {
			split := strings.Split(game, " ")
			return split[0], split[1]
		}()

		if me == "X" {
			current += 1
			if enemy == "A" {
				current += 3
			}
			if enemy == "C" {
				current += 6
			}
		} else if me == "Y" {
			current += 2
			if enemy == "B" {
				current += 3
			}
			if enemy == "A" {
				current += 6
			}
		} else { // Z
			current += 3
			if enemy == "C" {
				current += 3
			}
			if enemy == "B" {
				current += 6
			}
		}

		score += current
	}
	return score
}

func part2(input string) int {
	score := 0

	for _, game := range strings.Split(input, "\n") {
		current := 0
		enemy, me := func() (string, string) {
			split := strings.Split(game, " ")
			return split[0], split[1]
		}()

		if me == "X" {
			if enemy == "A" {
				current += 3
			} else if enemy == "B" {
				current += 1
			} else {
				current += 2
			}
		} else if me == "Y" {
			current += 3
			if enemy == "A" {
				current += 1
			} else if enemy == "B" {
				current += 2
			} else {
				current += 3
			}
		} else { // Z
			current += 6
			if enemy == "A" {
				current += 2
			} else if enemy == "B" {
				current += 3
			} else {
				current += 1
			}
		}

		score += current
	}
	return score
}
