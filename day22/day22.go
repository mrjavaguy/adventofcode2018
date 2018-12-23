package main

import (
	"container/heap"
	"fmt"
)

type Region struct {
	regionType    int
	erosionLevel  int
	geologicIndex int
}

const (
	Rocky  = 0
	Wet    = 1
	Narrow = 2
)

type Cave map[complex128]Region

var Neighbors = []complex128{-1i, -1, 1, 1i}

func main() {
	part1 := Day22Part1(8112, complex(13, 743))
	part2 := Day22Part2(8112, complex(13, 743))
	fmt.Printf("Day 22 Part 1 %v, Part 2 %v", part1, part2)
}

func makePoint(x, y int) complex128 {
	return complex(float64(x), float64(y))
}

func makeRegion(cave Cave, p complex128, target complex128, depth int) (region Region) {
	geologicIndex := 0
	x := int(real(p))
	y := int(imag(p))
	if p == 0 || p == target {
		geologicIndex = 0
	} else if y == 0 {
		geologicIndex = x * 16807
	} else if x == 0 {
		geologicIndex = y * 48271
	} else {
		r1, ok := cave[(p - 1)]
		if !ok {
			r1 = makeRegion(cave, (p - 1), target, depth)
		}
		r2, ok := cave[(p - 1i)]
		if !ok {
			r2 = makeRegion(cave, (p - 1i), target, depth)
		}
		geologicIndex = r1.erosionLevel * r2.erosionLevel
	}
	erosionLevel := (geologicIndex + depth) % 20183
	regionType := erosionLevel % 3
	if p == 0 || p == target {
		regionType = Rocky
	}
	cave[p] = Region{
		geologicIndex: geologicIndex,
		erosionLevel:  erosionLevel,
		regionType:    regionType,
	}
	return cave[p]
}

func Day22Part1(depth int, target complex128) int {
	yMax := int(imag(target) + 1)
	xMax := int(real(target) + 1)
	cave := Cave{}
	cost := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			p := makePoint(x, y)
			r := makeRegion(cave, p, target, depth)
			cost += r.regionType
		}
	}
	return cost
}

func Day22Part2(depth int, target complex128) int {
	cave := Cave{}
	makeRegion(cave, 0+0i, target, depth)

	result := Dijkstra(cave, target, depth)

	return result
}

func PrintRegionRune(regionType int) {
	switch regionType {
	case Rocky:
		fmt.Print(".")
	case Wet:
		fmt.Print("=")
	case Narrow:
		fmt.Print("|")
	}

}

const (
	None = 1 << iota
	Torch
	ClimbingGear
)

type ToolInUse struct {
	point complex128
	tool  int
	time  int
}

type PriorityQueue struct {
	items []ToolInUse
	// value to index
	m map[complex128]int
}

func (cave Cave) calcNeighbors(t ToolInUse, target complex128, depth int) []ToolInUse {
	neighbors := []ToolInUse{}
	p := t.point
	for _, n := range Neighbors {
		newP := p + n
		x := int(real(newP))
		y := int(imag(newP))
		if x < 0 || y < 0 || x > depth || y > depth {
			continue
		}
		r, ok := cave[newP]
		if !ok {
			r = makeRegion(cave, newP, target, depth)
		}
		if t.tool&allowed(r.regionType) != 0 {
			neighbors = append(neighbors, ToolInUse{point: newP, tool: t.tool, time: 1})
			neighbors = append(neighbors, ToolInUse{point: newP, tool: t.tool ^ allowed(r.regionType), time: 8})
		}

	}
	return neighbors
}

func allowed(regionType int) int {
	switch regionType {
	case Rocky:
		return ClimbingGear | Torch
	case Wet:
		return ClimbingGear | None
	case Narrow:
		return Torch | None
	default:
		panic(fmt.Errorf("unknown region type: %d", regionType))
	}
}

type equiqLoc struct {
	point complex128
	tool  int
}

func Dijkstra(cave Cave, target complex128, depth int) int {
	dist := make(map[equiqLoc]int)
	prev := make(map[equiqLoc]equiqLoc)
	t := equiqLoc{point: 0 + 0i, tool: Torch}
	dist[t] = 0
	q := &PriorityQueue{[]ToolInUse{}, make(map[complex128]int)}
	heap.Push(q, ToolInUse{point: 0 + 0i, tool: Torch})
	for len(q.items) != 0 {
		u := heap.Pop(q).(ToolInUse)
		uEquiploc := equiqLoc{point: u.point, tool: u.tool}
		if u.point == target && u.tool == Torch {
			current := uEquiploc
			for true {

				current = prev[current]
				if current.point == 0+0i {
					break
				}
			}
			return u.time
		}
		if t, ok := dist[uEquiploc]; ok && t < u.time {
			continue
		}
		for _, v := range cave.calcNeighbors(u, target, depth) {
			d := equiqLoc{point: v.point, tool: v.tool}
			t, ok := dist[d]
			alt := dist[uEquiploc] + v.time
			if !ok || alt < t {
				dist[d] = alt
				prev[d] = uEquiploc
				heap.Push(q, ToolInUse{point: v.point, time: alt, tool: v.tool})
			}
		}
	}

	return -1
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].time < pq.items[j].time
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i].point] = i
	pq.m[pq.items[j].point] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(ToolInUse)
	pq.m[item.point] = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item.point] = -1
	pq.items = old[0 : n-1]
	return item
}
