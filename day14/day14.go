package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 320851

func main() {
	recipes := []byte{'3', '7'}
	a, b := 0, 1

	for len(recipes) < 50000000 {
		score := []byte(strconv.Itoa(int(recipes[a] - '0' + recipes[b] - '0')))
		recipes = append(recipes, score...)

		a = (a + 1 + int(recipes[a]-'0')) % len(recipes)
		b = (b + 1 + int(recipes[b]-'0')) % len(recipes)
	}

	part1 := string(recipes[input : input+10])
	part2 := strings.Index(string(recipes), strconv.Itoa(input))

	fmt.Printf("Day 6 Part 1 %v, Part 2 %v", part1, part2)
}
