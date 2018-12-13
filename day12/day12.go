package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := "#...#####.#..##...##...#.##.#.##.###..##.##.#.#..#...###..####.#.....#..##..#.##......#####..####..."
	rules := []Rule{
		Rule{rule: "#.#.#", result: '#'},
		Rule{rule: "..###", result: '.'},
		Rule{rule: "#..#.", result: '#'},
		Rule{rule: ".#...", result: '#'},
		Rule{rule: "..##.", result: '#'},
		Rule{rule: "##.#.", result: '#'},
		Rule{rule: "##..#", result: '#'},
		Rule{rule: "####.", result: '#'},
		Rule{rule: "...#.", result: '#'},
		Rule{rule: "..#.#", result: '#'},
		Rule{rule: ".####", result: '#'},
		Rule{rule: "#.###", result: '.'},
		Rule{rule: "...##", result: '.'},
		Rule{rule: "..#..", result: '.'},
		Rule{rule: "#...#", result: '.'},
		Rule{rule: ".###.", result: '#'},
		Rule{rule: ".#.##", result: '.'},
		Rule{rule: ".##..", result: '#'},
		Rule{rule: "....#", result: '.'},
		Rule{rule: "#..##", result: '.'},
		Rule{rule: "##.##", result: '#'},
		Rule{rule: "#.##.", result: '.'},
		Rule{rule: "#....", result: '.'},
		Rule{rule: "##...", result: '#'},
		Rule{rule: ".#.#.", result: '.'},
		Rule{rule: "###.#", result: '#'},
		Rule{rule: "#####", result: '#'},
		Rule{rule: "#.#..", result: '.'},
		Rule{rule: ".....", result: '.'},
		Rule{rule: ".##.#", result: '.'},
		Rule{rule: "###..", result: '.'},
		Rule{rule: ".#..#", result: '.'},
	}
	gen := map[int]rune{}
	for i, s := range input {
		gen[i] = s
	}

	for i := 0; i < 200; i++ {
		gen = applyNextGeneration(gen, rules)
		if i == 20 {
			count := getCount(gen)
			fmt.Println(i, count)
		}
	}
	count := getCount(gen)
	result := ((50000000000/100)-2)*8600 + count
	fmt.Println(50000000000, result)
}

func nicePrint(input map[int]rune) {
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	min := keys[0]
	max := keys[len(keys)-1]
	for i := min; i <= max; i++ {
		if v, ok := input[i]; ok {
			fmt.Print(string(v))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func getCount(input map[int]rune) (count int) {
	count = 0
	for i, s := range input {
		if s == '#' {
			count += i
		}
	}

	return
}

type Rule struct {
	rule   string
	result rune
}

func applyRules(rules []Rule, input map[int]rune) map[int]rune {
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	min := keys[0]
	max := keys[len(keys)-1]

	newGen := map[int]rune{}

	for i := min - 2; i <= max+2; i++ {
		output := strings.Builder{}
		for j := -2; j < 3; j++ {
			if v, ok := input[i+j]; ok {
				output.WriteRune(v)
			} else {
				output.WriteRune('.')
			}
		}

		compare := output.String()

		for _, r := range rules {
			if r.rule == compare {
				newGen[i] = r.result
				break
			}
		}
	}

	return newGen
}

func applyNextGeneration(input map[int]rune, rules []Rule) map[int]rune {
	return applyRules(rules, input)
}
