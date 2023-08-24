package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

	for _, pair := range strings.Split(input, "\n") {
		elves := strings.Split(pair, ",")
		elf0 := strings.Split(elves[0], "-")
		elf1 := strings.Split(elves[1], "-")
		elf0_0, _ := strconv.Atoi(elf0[0])
		elf0_1, _ := strconv.Atoi(elf0[1])
		elf1_0, _ := strconv.Atoi(elf1[0])
		elf1_1, _ := strconv.Atoi(elf1[1])

		if (elf0_0 >= elf1_0 && elf0_1 <= elf1_1) || (elf1_0 >= elf0_0 && elf1_1 <= elf0_1) {
			result += 1
		}
	}

	return result
}

func part2(input string) int {
	result := 0

	for _, pair := range strings.Split(input, "\n") {
		elves := strings.Split(pair, ",")
		elf0 := strings.Split(elves[0], "-")
		elf1 := strings.Split(elves[1], "-")
		elf0_0, _ := strconv.Atoi(elf0[0])
		elf0_1, _ := strconv.Atoi(elf0[1])
		elf1_0, _ := strconv.Atoi(elf1[0])
		elf1_1, _ := strconv.Atoi(elf1[1])

		group0 := make(map[int]bool)

		for i := elf0_0; i <= elf0_1; i++ {
			group0[i] = true
		}
		group1 := make(map[int]bool)

		for i := elf1_0; i <= elf1_1; i++ {
			group1[i] = true
		}

		for x := range group0 {
			if group1[x] {
				result += 1
				break
			}
		}

	}

	return result
}
