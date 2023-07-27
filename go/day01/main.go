package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
	fmt.Println("Part 1:", ans1)

	ans2 := part2(input)
	fmt.Println("Part 2:", ans2)
}

func part1(input string) int {
	max := 0

	for _, entity := range strings.Split(input, "\n\n") {
		current := 0
		for _, calorie := range strings.Split(entity, "\n") {
			calorie, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			current += calorie
		}

		if current > max {
			max = current
		}
	}

	return max
}

func part2(input string) int {
	max_arr := []int{}

	for _, entity := range strings.Split(input, "\n\n") {
		current := 0
		for _, calorie := range strings.Split(entity, "\n") {
			calorie, err := strconv.Atoi(calorie)
			if err != nil {
				log.Fatal(err)
			}
			current += calorie
		}
		max_arr = append(max_arr, current)
	}

	sort.Ints(max_arr)
	sum := 0
	for _, num := range max_arr[len(max_arr)-3:] {
		sum += num
	}

	return sum
}
