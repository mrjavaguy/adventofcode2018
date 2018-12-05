package main

import (
	"adventofcode2018/input"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	guards := make(map[int]guard)

	lines, _ := input.FileToLines("day04/input04.txt")

	var parsedLines []inputLine
	for i := range lines {
		parsedLines = append(parsedLines, parseLine(lines[i]))
	}

	sort.Slice(parsedLines, func(i, j int) bool { return parsedLines[i].date.Before(parsedLines[j].date) })

	currentguard := -1
	for i := range parsedLines {
		currentguard = parseData(parsedLines[i], currentguard, guards)
	}

	part1 := day04Part1(guards)
	part2 := day04Part2(guards)
	fmt.Printf("Day 1 Part 1 %v, Part 2 %v", part1, part2)
}

func day04Part1(guards map[int]guard) int {
	maxGuard := -1
	maxTimeAsleep := -1
	for _, g := range guards {
		asleepTime := 0
		for _, i := range g.sleepIntervals {
			asleepTime += (i.end - i.start)
		}
		if asleepTime > maxTimeAsleep {
			maxTimeAsleep = asleepTime
			maxGuard = g.ID
		}
	}
	guard := guards[maxGuard]
	max := mostCommonMinute(guard.sleepIntervals)

	return guard.ID * max
}

func day04Part2(guards map[int]guard) int {
	minutes := make(map[int]map[int]int)
	for _, g := range guards {
		for _, i := range g.sleepIntervals {
			for j := i.start; j <= i.end; j++ {
				if minutes[j] == nil {
					minutes[j] = make(map[int]int)
				}
				minutes[j][g.ID]++
			}
		}
	}

	max := -1
	maxMinute := -1
	maxId := -1
	for i := 0; i < 60; i++ {
		if minutes[i] != nil {
			for j := range minutes[i] {
				if minutes[i][j] > max {
					max = minutes[i][j]
					maxMinute = i
					maxId = j
				}
			}
		}
	}

	return (maxMinute - 1) * maxId
}

func mostCommonMinute(intervals []interval) int {
	minutes := make([]int, 60)
	for _, i := range intervals {
		for j := i.start; j <= i.end; j++ {
			minutes[j] = minutes[j] + 1
		}
	}

	max := -1
	minute := -1
	for i, m := range minutes {
		if m > max {
			max = m
			minute = i - 1
		}
	}

	return minute
}

func parseDateTime(s string) time.Time {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println("error", err)
	}
	return t
}

type inputLine struct {
	date time.Time
	line string
}

type timeInterval struct {
	start time.Time
	end   time.Time
}

type interval struct {
	start int
	end   int
}

type guard struct {
	ID             int
	lastTimestamp  int
	asleep         bool
	sleepIntervals []interval
	awakeIntervals []interval
}

func parseLine(input string) inputLine {
	re := regexp.MustCompile("\\[(.*)\\] (.*)")
	matches := re.FindAllStringSubmatch(input, -1)
	return inputLine{
		date: parseDateTime(matches[0][1]),
		line: matches[0][2]}
}

func parseData(input inputLine, currentguard int, guards map[int]guard) int {
	if strings.Contains(input.line, "Guard") {
		if currentguard != -1 {
			guard := guards[currentguard]
			if guard.asleep {
				guard.sleepIntervals = append(guard.sleepIntervals, interval{start: guard.lastTimestamp + 1, end: 59})
			} else {
				guard.awakeIntervals = append(guard.awakeIntervals, interval{start: guard.lastTimestamp + 1, end: 59})
			}
		}
		if input.date.Hour() > 0 {
			input.date = input.date.AddDate(0, 0, 1)
			input.date = time.Date(input.date.Year(), input.date.Month(), input.date.Day(), 0, 0, 0, 0, input.date.Location())
		}
		re := regexp.MustCompile("#(\\d+)")
		matches := re.FindAllStringSubmatch(input.line, -1)
		currentguard, _ = strconv.Atoi(matches[0][1])
		if _, ok := guards[currentguard]; !ok {
			guards[currentguard] = guard{ID: currentguard, lastTimestamp: input.date.Minute()}
		}
	}

	if strings.Contains(input.line, "falls") {
		guard := guards[currentguard]
		if !guard.asleep {
			timestamp := input.date.Minute()
			guard.awakeIntervals = append(guard.awakeIntervals, interval{start: guard.lastTimestamp + 1, end: timestamp})
			guard.lastTimestamp = timestamp
			guard.asleep = true
		}
		guards[currentguard] = guard
	}

	if strings.Contains(input.line, "wakes") {
		guard := guards[currentguard]
		if guard.asleep {
			timestamp := input.date.Minute()
			guard.sleepIntervals = append(guard.sleepIntervals, interval{start: guard.lastTimestamp + 1, end: timestamp})
			guard.lastTimestamp = timestamp
			guard.asleep = false
		}
		guards[currentguard] = guard
	}
	return currentguard
}
