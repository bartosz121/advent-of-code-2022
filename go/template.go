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
	fmt.Println("Hello world")
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
