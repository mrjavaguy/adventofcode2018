package main

import (
	"adventofcode2018/input"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	lines, _ := input.FileToLines("day07/input07.txt")

	part1 := Day07Part1(lines)
	part2 := Day07Part2(lines)
	fmt.Printf("Day 7 Part 1 %v, Part 2 %v", part1, part2)
}

func Day07Part1(lines []string) string {
	steps := parseLines(lines)

	done := make([]string, 0)
	for step, parents := range steps {
		if len(parents) == 0 {
			done = append(done, step)
		}
	}

	s := strings.Builder{}
	for len(done) > 0 {
		sort.Slice(done, func(i, j int) bool { return done[i] < done[j] })
		first := done[0]
		done = append(done[:0], done[1:]...)
		s.WriteString(first)
		for step, parents := range steps {
			if i := stringInSlice(first, parents); i != -1 {
				steps[step] = append(parents[:i], parents[i+1:]...)
				if len(steps[step]) == 0 {
					done = append(done, step)
				}
			}
		}
	}

	return s.String()
}

func Day07Part2(lines []string) int {
	steps := parseLines(lines)

	ready := make([]string, 0)
	for step, parents := range steps {
		if len(parents) == 0 {
			ready = append(ready, step)
		}
	}

	workers := make([]worker, 5)

	timeStep := 0
	working := 1

	for working > 0 {
		working = 0
		for i, w := range workers {
			if w.timeLeft > 0 {
				w.timeLeft--
				workers[i] = w
				working++
			} else {
				if w.task != "" {
					finishedTask := w.task
					w.task = ""
					workers[i] = w
					for step, parents := range steps {
						if i := stringInSlice(finishedTask, parents); i != -1 {
							steps[step] = append(parents[:i], parents[i+1:]...)
							if len(steps[step]) == 0 {
								ready = append(ready, step)
							}
						}
					}
				}
			}
		}

		for len(ready) > 0 && working < len(workers) {
			sort.Slice(ready, func(i, j int) bool { return ready[i] < ready[j] })
			first := ready[0]
			ready = append(ready[:0], ready[1:]...)
			for i, w := range workers {
				if w.task == "" {
					w.task = first
					w.timeLeft = int(first[0]) - 'A' + 60
					workers[i] = w
					working++
					break
				}
			}
		}

		timeStep++
	}

	timeStep--

	return timeStep
}

type worker struct {
	task     string
	timeLeft int
}

func parseLines(lines []string) map[string][]string {
	steps := make(map[string][]string)
	re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	for _, s := range lines {
		matches := re.FindAllStringSubmatch(s, -1)
		if _, ok := steps[matches[0][2]]; !ok {
			steps[matches[0][2]] = make([]string, 0)
		}
		if _, ok := steps[matches[0][1]]; !ok {
			steps[matches[0][1]] = make([]string, 0)
		}
		steps[matches[0][2]] = append(steps[matches[0][2]], matches[0][1])
	}
	return steps
}

func stringInSlice(a string, list []string) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}
