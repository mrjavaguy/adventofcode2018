package main

import (
	"adventofcode2018/input"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day23/input23.txt")
	swarm, masterBot := ParseInput(lines)

	part1 := Day23Part1(swarm, masterBot)
	part2 := Day23Part2(swarm)
	fmt.Printf("Day 23 Part 1 %v, Part 2 %v", part1, part2)
}

func Day23Part1(swarm Swarm, masterBot Nanobot) int {
	c := 0
	for _, b := range swarm {
		if masterBot.inRange(b) {
			c++
		}
	}
	return c
}

func Day23Part2(swarm Swarm) int {
	xMap := map[int]int{}
	for _, bot := range swarm {
		xMin := bot.location.x + bot.location.y + bot.location.z - bot.powerRange
		xMax := bot.location.x + bot.location.y + bot.location.z + bot.powerRange + 1
		xMap[xMin] += 1
		xMap[xMax] -= 1
	}

	keys := []int{}
	for k := range xMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	running := 0
	max := 0
	maxStart := 0
	maxEnd := 0

	for _, k := range keys {
		running += xMap[k]
		if running > max {
			max = running
			maxStart = k
		}
	}

	for _, k := range keys {
		if k > maxStart {
			maxEnd = k
			break
		}
	}
	return maxEnd - 1
}

type Candidate struct {
	Loc     Coordinate
	Quality int
}

type bronKerboschState struct {
	neighbors  map[int][]int
	maxCliques [][]int
}

func BronKerbosch(bronKerboschState *bronKerboschState, r, p, x intSet) {
	if len(p) == 0 && len(x) == 0 {
		if len(bronKerboschState.maxCliques) > 0 && len(bronKerboschState.maxCliques[0]) > len(r) {
			return
		}
		if len(bronKerboschState.maxCliques) > 0 && len(bronKerboschState.maxCliques[0]) < len(r) {
			// Found a longer clique.
			bronKerboschState.maxCliques = nil
		}
		clique := make([]int, 0, len(r))
		for v := range r {
			clique = append(clique, v)
		}
		sort.Ints(clique)
		bronKerboschState.maxCliques = append(bronKerboschState.maxCliques, clique)
		return
	}
	fmt.Print(".")
	u := -1
	if len(p) > 0 {
		for v := range p {
			u = v
			break
		}
	} else {
		for v := range x {
			u = v
			break
		}
	}
	nu := bronKerboschState.neighbors[u]
	nuSet := make(intSet, len(nu))
	for _, uu := range nu {
		nuSet.add(uu)
	}
	for v := range p {
		if nuSet.contains(v) {
			continue
		}
		ns := bronKerboschState.neighbors[v]
		p1 := make(intSet, len(ns))
		x1 := make(intSet, len(ns))
		for _, n := range ns {
			if p.contains(n) {
				p1.add(n)
			}
			if x.contains(n) {
				x1.add(n)
			}
		}
		r.add(v)
		BronKerbosch(bronKerboschState, r, p1, x1)
		r.remove(v)
		p.remove(v)
		x.add(v)
	}
}

type Coordinate struct {
	x, y, z int
}

type Nanobot struct {
	location   Coordinate
	powerRange int
}

func (bot1 *Nanobot) distance(bot2 Nanobot) int {
	return int(math.Abs(float64(bot1.location.x)-float64(bot2.location.x)) + math.Abs(float64(bot1.location.y)-float64(bot2.location.y)) + math.Abs(float64(bot1.location.z)-float64(bot2.location.z)))
}

func (bot1 *Nanobot) distanceToRange(bot2 Nanobot) int {
	return bot1.distance(bot2) - bot1.powerRange
}

func (bot1 *Nanobot) inRange(bot2 Nanobot) bool {
	return bot1.distance(bot2) <= bot1.powerRange
}

func (bot1 *Nanobot) getAllNeighbors(swarm Swarm) (neighbors []int) {
	neighbors = []int{}
	for j, b := range swarm {
		if *bot1 == b {
			continue
		}
		d := int(bot1.distance(b))
		if d <= bot1.powerRange && d <= b.powerRange {
			neighbors = append(neighbors, j)
		}
	}
	return
}

func (current *Coordinate) WalkPartway(target Coordinate, distance int) Coordinate {
	xDelta := target.x - current.z
	yDelta := target.y - current.y
	zDelta := target.z - current.z

	sum := 0
	if xDelta < 0 {
		sum -= xDelta
	} else {
		sum += xDelta
	}
	if yDelta < 0 {
		sum -= yDelta
	} else {
		sum += yDelta
	}
	if zDelta < 0 {
		sum -= zDelta
	} else {
		sum += zDelta
	}

	return Coordinate{current.x + distance*xDelta/sum, current.y + distance*yDelta/sum, current.z + distance*zDelta/sum}
}

type Swarm []Nanobot

func ParseInput(lines []string) (swarm Swarm, biggestRadius Nanobot) {
	re := regexp.MustCompile("(-?\\d+)")
	swarm = Swarm{}
	br := 0
	for _, l := range lines {
		matches := re.FindAllStringSubmatch(l, -1)
		x, _ := strconv.Atoi(matches[0][0])
		y, _ := strconv.Atoi(matches[1][0])
		z, _ := strconv.Atoi(matches[2][0])
		r, _ := strconv.Atoi(matches[3][0])
		bot := Nanobot{
			location:   Coordinate{x, y, z},
			powerRange: r,
		}
		if r > br {
			biggestRadius = bot
			br = r
		}
		swarm = append(swarm, bot)
	}
	return
}

type intSet map[int]struct{}

func (s intSet) add(i int)    { s[i] = struct{}{} }
func (s intSet) remove(i int) { delete(s, i) }
func (s intSet) copy() intSet {
	newSet := make(intSet, len(s))
	for i := range s {
		newSet[i] = struct{}{}
	}
	return newSet
}
func (s intSet) contains(i int) bool {
	_, ok := s[i]
	return ok
}
