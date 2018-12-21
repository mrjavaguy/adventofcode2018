package main

import (
	"adventofcode2018/input"
	"container/heap"
	"fmt"
	"math"
)

func main() {
	line, _ := input.FileToLine("day20/input20.txt")

	part1 := Day20Part1(line)
	part2 := Day20Part2(line)
	fmt.Printf("Day 20 Part 1 %v, Part 2 %v", part1, part2)
}

func Day20Part1(line string) int {
	base, xMin, xMax, yMin, yMax := BuildMap(line)
	PrintMap(base, xMin, xMax, yMin, yMax)
	dist, _ := Dijkstra(base)
	x := 0
	for _, d := range dist {
		if d > x {
			x = d
		}
	}
	return x
}

func Day20Part2(line string) int {
	base, _, _, _, _ := BuildMap(line)
	dist, _ := Dijkstra(base)
	x := 0
	for _, d := range dist {
		if d >= 1000 {
			x++
		}
	}
	return x
}

func makeroom(base Map, x, y int) Room {
	p := makePoint(x, y)
	if _, ok := base[p]; !ok {
		base[p] = Room{}
	}
	return base[p]
}

func BuildMap(line string) (base Map, xMin int, xMax int, yMin int, yMax int) {
	stack := NewStack()
	base = Map{}
	x, y := 0, 0
	xMin, xMax, yMin, yMax = 0, 0, 0, 0
	for _, c := range line {
		p := makePoint(x, y)
		r := makeroom(base, x, y)
		switch c {
		case 'N':
			r.WallType[North] = Door
			base[p] = r
			y--
			r = makeroom(base, x, y)
			r.WallType[South] = Door
			p := makePoint(x, y)
			base[p] = r
			if y < yMin {
				yMin = y
			}
		case 'S':
			r.WallType[South] = Door
			base[p] = r
			y++
			r = makeroom(base, x, y)
			r.WallType[North] = Door
			p := makePoint(x, y)
			base[p] = r
			if y > yMax {
				yMax = y
			}
		case 'W':
			r.WallType[West] = Door
			base[p] = r
			x--
			r = makeroom(base, x, y)
			r.WallType[East] = Door
			p := makePoint(x, y)
			base[p] = r
			if x < xMin {
				xMin = x
			}
		case 'E':
			r := base[p]
			r.WallType[East] = Door
			base[p] = r
			x++
			r = makeroom(base, x, y)
			r.WallType[West] = Door
			p := makePoint(x, y)
			base[p] = r
			if x > xMax {
				xMax = x
			}
		case '(':
			stack.Push(&p)
		case '|':
			p := stack.Pop()
			x = int(real(*p))
			y = int(imag(*p))
			stack.Push(p)
		case ')':
			stack.Pop()
		}
	}
	return
}

func makePoint(x, y int) complex128 {
	return complex(float64(x), float64(y))
}

func PrintMap(base Map, xMin int, xMax int, yMin int, yMax int) {
	for y := yMin; y <= yMax; y++ {
		for i := 0; i < 2; i++ {
			for x := xMin; x <= xMax; x++ {
				p := makePoint(x, y)
				r := base[p]
				c := "#"
				if i == 0 {
					if r.WallType[North] == Door {
						c = "-"
					}
					fmt.Print("#")
					fmt.Print(c)
				} else {
					if r.WallType[West] == Door {
						c = "|"
					}
					fmt.Print(c)
					if x == 0 && y == 0 {
						fmt.Print("X")
					} else {
						fmt.Print(".")
					}
				}
			}
			fmt.Println("#")
		}
	}
	for x := xMin; x <= xMax; x++ {
		fmt.Print("##")
	}
	fmt.Println("#")
}

type Map map[complex128]Room

type Room struct {
	WallType [4]int
}

const (
	North = 0
	West  = 1
	East  = 2
	South = 3
)

const (
	Unknown = 0
	Wall    = 1
	Door    = 2
)

func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	items []*complex128
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *complex128) {
	s.items = append(s.items[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *complex128 {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.items[s.count]
}

type PriorityQueue struct {
	items []complex128
	// value to index
	m map[complex128]int
	// value to priority
	pr map[complex128]int
}

func Dijkstra(base Map) (dist map[complex128]int, prev map[complex128]complex128) {
	dist = make(map[complex128]int)
	prev = make(map[complex128]complex128)
	sid := 0 + 0i
	dist[sid] = 0
	q := &PriorityQueue{[]complex128{}, make(map[complex128]int), make(map[complex128]int)}
	vertices := []complex128{sid}
	for p := range base {
		vertices = append(vertices, p)
	}
	for _, v := range vertices {
		if v != sid {
			dist[v] = math.MaxInt32
		}
		prev[v] = complex(math.MaxInt32, math.MaxUint32)
		q.addWithPriority(v, dist[v])
	}
	var p complex128
	for len(q.items) != 0 {
		u := heap.Pop(q).(complex128)

		neighbors := []complex128{}
		for x := 0; x < 4; x++ {
			switch x {
			case 0:
				p = u - 1i
			case 1:
				p = u - 1
			case 2:
				p = u + 1
			case 3:
				p = u + 1i
			}
			if base[u].WallType[x] == Door {
				neighbors = append(neighbors, p)
			}
		}
		for _, v := range neighbors {
			alt := dist[u] + 1
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				q.update(v, alt)
			}
		}
	}

	return
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool {
	if pq.pr[pq.items[i]] == pq.pr[pq.items[j]] {
		if imag(pq.items[i]) == imag(pq.items[j]) {
			return real(pq.items[i]) < real(pq.items[j])
		}
		return imag(pq.items[i]) < imag(pq.items[j])
	}
	return pq.pr[pq.items[i]] < pq.pr[pq.items[j]]
}
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(complex128)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

// update modifies the priority of an item in the queue.
func (pq *PriorityQueue) update(item complex128, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *PriorityQueue) addWithPriority(item complex128, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}
