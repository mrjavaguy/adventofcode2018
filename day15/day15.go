package main

import (
	"adventofcode2018/input"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	args := os.Args
	inputFile := "day15/input15.txt"
	if len(args) > 1 {
		inputFile = "day15/" + args[1]
	}
	lines, _ := input.FileToLines(inputFile)

	part1 := Day15Part1(lines)
	part2 := Day15Part2(lines)
	fmt.Printf("Day 1 Part 1 %v, Part 2 %v", part1, part2)
}

func Day15Part1(input []string) int {
	cave := NewCave(input, 3)
	//cave.printCave()
	return Solve(cave, false)
	return -1
}

func Solve(cave *Cave, stopOnElfDeath bool) int {
	for round := 0; true; round++ {
		hitPointsLeft, combatHappened := cave.CurrentStatus()
		if !combatHappened {
			//fmt.Println("After", round+1, ":")
			//cave.printCave()
			return round * hitPointsLeft
		}

		if cleanRound := cave.DoRound(stopOnElfDeath); !cleanRound {
			round--
		}

		//fmt.Println("After", round+1, ":")
	}
	return -1
}

func Day15Part2(input []string) int {
	for x := 4; true; x++ {
		cave := NewCave(input, x)
		result := Solve(cave, true)
		fmt.Println(x, result)
		onlyElves := true
		for _, c := range cave.Combatants {
			if (c.class == 'G' && c.hitPoints > 0) || (c.class == 'E' && c.hitPoints < 0) {
				onlyElves = false
				break
			}
		}
		if onlyElves {
			cave.printCave()
			return result
		}
	}
	return -1
}

type Point struct {
	x, y int
}

type Cave struct {
	Combatants Combatants
	Map        Map
	Size       Point
}

func NewCave(input []string, elfPower int) *Cave {
	c := &Cave{}
	c.ParseMap(input, elfPower)
	return c
}

func (cave *Cave) ParseMap(input []string, elfPower int) {
	m := make(Map)

	for y, row := range input {
		for x, col := range row {

			tile := &Tile{kind: col}
			if col == 'E' {
				cave.Combatants = append(cave.Combatants, NewCombatant(tile, col, elfPower))
			}
			if col == 'G' {
				cave.Combatants = append(cave.Combatants, NewCombatant(tile, col, 3))
			}

			m.SetTile(tile, Point{x, y})
		}
	}
	cave.Map = m
	cave.Size = Point{len(input[0]), len(input)}
}

func (cave *Cave) printCave() {
	for y := 0; y < cave.Size.y; y++ {
		var combatants []string
		for x := 0; x < cave.Size.x; x++ {
			t := cave.Map[Point{x, y}]
			if t == nil {
				continue
			}
			fmt.Print(string(t.kind))

			if t.combatant != nil {
				combatants = append(combatants, fmt.Sprintf("%c(%d)", t.combatant.class, t.combatant.hitPoints))
			}
		}
		if len(combatants) > 0 {
			fmt.Print("  ", strings.Join(combatants, ", "))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (cave *Cave) removeCombatant(c *Combatant) {
	cave.Map[c.position.location].kind = '.'
	c.position.kind = '.'
	c.position.combatant = nil
	c.position = nil
}

func (cave *Cave) CurrentStatus() (int, bool) {
	var elves, goblins bool
	var hp int

	for _, c := range cave.Combatants {
		if c.hitPoints <= 0 {
			continue
		}
		if c.class == 'E' {
			elves = true
		} else {
			goblins = true
		}
		hp += c.hitPoints
	}

	return hp, elves && goblins
}

func (cave *Cave) RemoveTheDead() {
	var newCombatants Combatants
	for _, combatant := range cave.Combatants {
		if combatant.hitPoints > 0 {
			newCombatants = append(newCombatants, combatant)
		}
	}
	cave.Combatants = newCombatants
}

func (cave *Cave) DoRound(stopOnElfDeath bool) bool {
	sort.Sort(cave.Combatants)
	for _, combatant := range cave.Combatants {
		if combatant.hitPoints <= 0 {
			continue
		}
		if !combatant.Targets(cave) {
			return false
		}
		combatant.Move(cave)
		if combatant.Attack(cave) && stopOnElfDeath {
			return false
		}
	}
	return true
}

type Combatant struct {
	position    *Tile
	hitPoints   int
	attackPower int
	class       rune
}

type Combatants []*Combatant

func (c Combatants) Len() int      { return len(c) }
func (c Combatants) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Combatants) Less(i, j int) bool {
	if c[i].hitPoints <= 0 {
		return false
	}
	if c[i].position == nil {
		return false
	}
	if c[j].position == nil {
		return true
	}
	if c[i].position.location.y == c[j].position.location.y {
		return c[i].position.location.x < c[j].position.location.x
	}
	return c[i].position.location.y < c[j].position.location.y
}

func NewCombatant(tile *Tile, class rune, power int) *Combatant {
	combatant := &Combatant{
		class:       class,
		hitPoints:   200,
		attackPower: power,
		position:    tile,
	}
	tile.combatant = combatant

	return combatant
}

func (c *Combatant) Targets(cave *Cave) bool {
	for _, combatant := range cave.Combatants {
		if combatant.class != c.class && combatant.hitPoints > 0 {
			return true
		}
	}
	return false
}

func (c *Combatant) TakeDamage(cave *Cave, attackPower int) bool {
	c.hitPoints -= attackPower
	if c.hitPoints <= 0 {
		cave.removeCombatant(c)
		return true
	}
	return false
}

func (c *Combatant) Attack(cave *Cave) bool {
	enemy := c.EnemyNeighbor(cave)
	if enemy != nil {
		killed := enemy.TakeDamage(cave, c.attackPower)
		return killed && enemy.class == 'E'
	}
	return false
}

func (c *Combatant) Enemies(cave *Cave) Combatants {
	var enemies Combatants
	for _, combatant := range cave.Combatants {
		if combatant.class != c.class && combatant.hitPoints > 0 {
			enemies = append(enemies, combatant)
		}
	}
	sort.Sort(enemies)
	return enemies
}

func (c *Combatant) EnemyNeighbor(cave *Cave) *Combatant {
	var target *Combatant
	for _, offset := range readOrderOffset {
		t := cave.Map[c.position.location.Add(offset)]
		if t != nil && t.combatant != nil && t.combatant.class != c.class && t.combatant.hitPoints > 0 {
			if target == nil || t.combatant.hitPoints < target.hitPoints {
				target = t.combatant
			}
		}
	}
	return target
}

func (c *Combatant) Move(cave *Cave) {
	if c.EnemyNeighbor(cave) != nil {
		return
	}
	next := c.NextTile(cave)
	if next != nil {
		next.combatant = c
		next.kind = c.class
		c.position.kind = '.'
		c.position.combatant = nil
		c.position = next
	}
}

func (c *Combatant) NextTile(cave *Cave) *Tile {
	targets := SortableTiles{}

	closestTargetDistance := math.MaxInt32
	distances, path := c.Dijkstra(cave)
	enemies := c.Enemies(cave)
	for _, enemy := range enemies {
		for _, target := range enemy.position.WalkableNeighbors(cave) {
			if distance := distances[target.location]; distance < math.MaxInt32 && distance <= closestTargetDistance {
				if distance < closestTargetDistance {
					closestTargetDistance = distance
					targets = SortableTiles{}
				}
				targets = append(targets, target)
			}
		}
	}
	sort.Sort(targets)
	if len(targets) > 0 {
		current := targets[0]
		for {
			if path[current.location] == c.position.location {
				return current
			}
			current = cave.Map[path[current.location]]
		}
	}
	return nil

}

func (c *Combatant) Dijkstra(cave *Cave) (dist map[Point]int, prev map[Point]Point) {
	dist = make(map[Point]int)
	prev = make(map[Point]Point)
	sid := c.position.location
	dist[sid] = 0
	q := &PriorityQueue{[]Point{}, make(map[Point]int), make(map[Point]int)}
	vertices := []Point{sid}
	for p, t := range cave.Map {
		if t.kind == '.' {
			vertices = append(vertices, p)
		}
	}
	for _, v := range vertices {
		if v != sid {
			dist[v] = math.MaxInt32
		}
		prev[v] = Point{-1, -1}
		q.addWithPriority(v, dist[v])
	}
	var p Point
	for len(q.items) != 0 {
		u := heap.Pop(q).(Point)

		neighbors := []Point{}
		for _, x := range readOrderOffset {
			p = u.Add(x)
			t := cave.Map[p]
			if t.kind == '.' {
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

var readOrderOffset = []Point{
	{0, -1},
	{-1, 0},
	{1, 0},
	{0, 1},
}

func (p Point) Add(vector Point) Point {
	return Point{p.x + vector.x, p.y + vector.y}
}

type Map map[Point]*Tile

func (m Map) SetTile(tile *Tile, p Point) {
	m[p] = tile
	tile.location = p
}

type Tile struct {
	location  Point
	kind      rune
	combatant *Combatant
}

func (t Tile) WalkableNeighbors(cave *Cave) []*Tile {
	var neighbors []*Tile

	for _, offset := range readOrderOffset {
		if n := cave.Map[t.location.Add(offset)]; n != nil && n.kind == '.' {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

type SortableTiles []*Tile

func (s SortableTiles) Len() int {
	return len(s)
}

func (s SortableTiles) Less(i, j int) bool {
	if s[i].location.y == s[j].location.y {
		return s[i].location.x < s[j].location.x
	}
	return s[i].location.y < s[j].location.y
}

func (s SortableTiles) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type PriorityQueue struct {
	items []Point
	// value to index
	m map[Point]int
	// value to priority
	pr map[Point]int
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool {
	if pq.pr[pq.items[i]] == pq.pr[pq.items[j]] {
		if pq.items[i].y == pq.items[j].y {
			return pq.items[i].x < pq.items[j].x
		}
		return pq.items[i].y < pq.items[j].y
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
	item := x.(Point)
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
func (pq *PriorityQueue) update(item Point, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *PriorityQueue) addWithPriority(item Point, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}
